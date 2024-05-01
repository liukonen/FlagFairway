package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgraph-io/badger/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/robfig/cron/v3"
)

var (
	db *badger.DB
	cronRunner *cron.Cron
)
// @title           Flag Fairway
// @version         0.1
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   Luke Liukonen
// @contact.url    https://liukonen.dev
// @BasePath  /api/v1
func main() {

	var err error
	// Open Badger database
	db, err = badger.Open(badger.DefaultOptions("./data"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	// Run garbage collection in a separate Goroutine
	location, _ := time.LoadLocation("America/Chicago")
	cronRunner := cron.New(cron.WithLocation(location))

	cronRunner.AddJob("0 0,6,12,18 * * *", cron.FuncJob(func() {
		err := db.RunValueLogGC(0.7)
		if err != nil {
			log.Printf("Error running garbage collection: %v", err)
		}
	}))
	cronRunner.Start()

	app := fiber.New()
	app.Static("/", "./internal/ui/build")

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:          "/swagger2/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))
	app.Static("/swagger2/doc.json", "./docs/swagger.json")
	app.Get("/api/v1/feature_flags", getFeatureFlags)
	app.Get("/api/v1/feature_flags/:key", getFeatureFlag)
	app.Post("/api/v1/feature_flags/:key", CreateFeatureFlag)
	app.Put("/api/v1/feature_flags/:key", createOrUpdateFeatureFlag)
	app.Delete("/api/v1/feature_flags/:key", deleteFeatureFlag)
	app.Get("/api/v1/health", getHealth)
	log.Println("Server started on port 8080")
	log.Fatal(app.Listen(":8080"))
}

// getFeatureFlags godoc
// @Summary      get Feature Flags
// @Tags feature_flags
// @Description  get list of current feature flags
// @Produce      json
// @Success      200  {object}  []string
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /feature_flags [get]
func getFeatureFlags(c *fiber.Ctx) error {
	var featureFlags []string
	err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			featureFlags = append(featureFlags, string(it.Item().Key()))
		}

		return nil
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(featureFlags)
}

func createOrUpdateFeatureFlag(c *fiber.Ctx) error {
	key := c.Params("key")

	body := c.Body()
	fmt.Print(key, body)
	err := addOrUpdateFlag(key, string(body))
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(http.StatusAccepted)
}

// @Summary Update a feature flag
// @Description Create a new feature flag if it doesn't exist or update an existing one
// @ID create-or-update-feature-flag
// @Tags feature_flags
// @Param key path string true "Key of the feature flag"
// @Param body body string true "New value of the feature flag"
// @Success 202 {string} string "Feature flag created or updated"
// @Failure 409 {string} string "Invalid request body"
// @Failure 500 {string} string "Internal server error"
// @Router /feature_flags/{key} [put]
func UpdateFeatureFlag(c *fiber.Ctx) error {
	//Check if feature flag exists, if so, update, else error
	key := c.Params("key")
	_, err := getFlag(key)
	if err != nil && err.Error() == "Key not found" {
		return c.Status(http.StatusConflict).SendString("Feature flag not found")
	}
	return createOrUpdateFeatureFlag(c)
}

// @Summary Create a feature flag
// @Description Create a new feature flag if it doesn't exist or update an existing one
// @ID create-or-update-feature-flag
// @Tags feature_flags
// @Param key path string true "Key of the feature flag"
// @Param body body string true "New value of the feature flag"
// @Success 202 {string} string "Feature flag created or updated"
// @Failure 409 {string} string "Invalid request body"
// @Failure 500 {string} string "Internal server error"
// @Router /feature_flags/{key} [post]
func CreateFeatureFlag(c *fiber.Ctx) error {
	key := c.Params("key")
	_, err := getFlag(key)
	if err == nil {
		return c.Status(http.StatusConflict).SendString("Feature flag found")
	}
	return createOrUpdateFeatureFlag(c)
}

// @Summary Delete a feature flag by key
// @Description Delete a feature flag by its key
// @ID delete-feature-flag
// @Tags feature_flags
// @Param key path string true "Key of the feature flag to delete"
// @Success 200 {string} string "Feature flag deleted"
// @Failure 404 {string} string "Feature flag not found"
// @Failure 500 {string} string "Internal server error"
// @Router /feature_flags/{key} [delete]
func deleteFeatureFlag(c *fiber.Ctx) error {
	key := c.Params("key") //strings.TrimPrefix(r.URL.Path, "/api/v1/feature_flags/")
	fmt.Print(key)
	err := deleteFlag(key)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(http.StatusOK)
}

// @Summary Get a feature flag by key
// @Description Retrieve the value of a feature flag by its key
// @ID get-feature-flag
// @Tags feature_flags
// @Param key path string true "Key of the feature flag"
// @Success 200 {string} string "Value of the feature flag"
// @Failure 404 {string} string "Feature flag not found"
// @Failure 500 {string} string "Internal server error"
// @Router /feature_flags/{key} [get]
func getFeatureFlag(c *fiber.Ctx) error {
	key := c.Params("key") //strings.TrimPrefix(r.URL.Path, "/api/v1/feature_flags/")
	fmt.Print(key)
	flag, err := getFlag(key)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString(flag)
}

// @Summary Get the health status of the application
// @Description Returns the health status of the application
// @ID get-health
// @Success 200 {string} string "Healthy"
// @Router /api/v1/health [get]
func getHealth(c *fiber.Ctx)error {
	return c.Status(http.StatusOK).SendString("Healthy")
}
func getFlag(key string) (string, error) {
	var flagValue string

	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			flagValue = string(val)
			return nil
		})
	})
	if err != nil {
		return "", err
	}

	return flagValue, nil
}

// Function to create or update the value of a feature flag in the database
func addOrUpdateFlag(key, value string) error {
	fmt.Print(key, value)
	return db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(value))
	})
}

func deleteFlag(key string) error {
	fmt.Print(key)
	return db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}
