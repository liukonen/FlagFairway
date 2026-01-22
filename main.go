package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"io"
	"github.com/dgraph-io/badger/v4"
	"github.com/robfig/cron/v3"
	"github.com/labstack/echo/v5"
)

var (
	db *badger.DB
	cronRunner *cron.Cron
)

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
	cronRunner = cron.New(cron.WithLocation(location))

	cronRunner.AddJob("0 0,6,12,18 * * *", cron.FuncJob(func() {
		err := db.RunValueLogGC(0.7)
		if err != nil {
			log.Printf("Error running garbage collection: %v", err)
		}
	}))
	cronRunner.Start()

	app := echo.New()//fiber.New()
	app.Static("/", "./internal/ui/build")

	app.GET("/api/v1/feature_flags", getFeatureFlags)
	app.GET("/api/v1/feature_flags/:key", getFeatureFlag)
	app.POST("/api/v1/feature_flags/:key", CreateFeatureFlag)
	app.PUT("/api/v1/feature_flags/:key", UpdateFeatureFlag)
	app.DELETE("/api/v1/feature_flags/:key", deleteFeatureFlag)
	app.GET("/api/v1/health", getHealth)
	log.Println("Server started on port 8080")
	log.Fatal(app.Start(":8080"))
}

// @Summary      get Feature Flags
// @Tags feature_flags
// @Description  get list of current feature flags
// @Produce      json
// @Success      200  {object}  []string
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /feature_flags [get]
func getFeatureFlags(c *echo.Context) error {
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
		return c.String(http.StatusServiceUnavailable, err.Error())
	}
	return c.JSON(http.StatusOK, featureFlags)
}

func createOrUpdateFeatureFlag(c *echo.Context) error {
	key := c.Param("key")
	body, _ := RequestBody(c)
	fmt.Print(key, body)
	err := addOrUpdateFlag(key, string(body))
	if err != nil {
		return c.String(http.StatusBadRequest,err.Error())
	}
	return c.String(http.StatusAccepted, "")
}

// @Description Create a new feature flag if it doesn't exist or update an existing one
// @ID create-or-update-feature-flag
// @Tags feature_flags
// @Param key path string true "Key of the feature flag"
// @Param body body string true "New value of the feature flag"
// @Success 202 {string} string "Feature flag created or updated"
// @Failure 409 {string} string "Invalid request body"
// @Failure 500 {string} string "Internal server error"
// @Router /feature_flags/{key} [put]
func UpdateFeatureFlag(c *echo.Context) error {
	key := c.Param("key")
	_, err := getFlag(key)
	if err != nil && err.Error() == "Key not found" {
		return c.String(http.StatusConflict,"Feature flag not found")
	}
	return createOrUpdateFeatureFlag(c)
}

// @Description Create a new feature flag if it doesn't exist or update an existing one
// @ID create-or-update-feature-flag
// @Tags feature_flags
// @Param key path string true "Key of the feature flag"
// @Param body body string true "New value of the feature flag"
// @Success 202 {string} string "Feature flag created or updated"
// @Failure 409 {string} string "Invalid request body"
// @Failure 500 {string} string "Internal server error"
// @Router /feature_flags/{key} [post]
func CreateFeatureFlag(c *echo.Context) error {
	key := c.Param("key")
	_, err := getFlag(key)
	if err == nil {
		return c.String(http.StatusConflict,"Feature flag found")
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
func deleteFeatureFlag(c *echo.Context) error {
	key := c.Param("key") 
	fmt.Print(key)
	err := deleteFlag(key)
	if err != nil {
		return c.String(http.StatusInternalServerError,err.Error())
	}
	return c.String(http.StatusOK,"")
}


func getFeatureFlag(c *echo.Context) error {
	key := c.Param("key") 
	fmt.Print(key)
	flag, err := getFlag(key)
	if err != nil {
		return c.String(http.StatusNotFound,err.Error())
	}
	return c.String(http.StatusOK,flag)
}


// @Description Returns the health status of the application
// @ID get-health
// @Success 200 {string} string "Healthy"
// @Router /api/v1/health [get]
func getHealth(c *echo.Context)error {
	return c.String(http.StatusOK, "Healthy")
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

func RequestBody(c *echo.Context) (string, error){
	bodyBytes, err := io.ReadAll(c.Request().Body)
        if err != nil {
            return "", err
        }
        return string(bodyBytes), nil
}
