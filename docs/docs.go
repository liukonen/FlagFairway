// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Luke Liukonen",
            "url": "https://liukonen.dev"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/feature_flags": {
            "get": {
                "description": "get list of current feature flags",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "feature_flags"
                ],
                "summary": "get Feature Flags",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/feature_flags/{key}": {
            "get": {
                "description": "Retrieve the value of a feature flag by its key",
                "tags": [
                    "feature_flags"
                ],
                "summary": "Get a feature flag by key",
                "operationId": "get-feature-flag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Key of the feature flag",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Value of the feature flag",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Feature flag not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new feature flag if it doesn't exist or update an existing one",
                "tags": [
                    "feature_flags"
                ],
                "summary": "Create or update a feature flag",
                "operationId": "create-or-update-feature-flag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Key of the feature flag",
                        "name": "key",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "New value of the feature flag",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Feature flag created or updated",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a feature flag by its key",
                "tags": [
                    "feature_flags"
                ],
                "summary": "Delete a feature flag by key",
                "operationId": "delete-feature-flag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Key of the feature flag to delete",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Feature flag deleted",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Feature flag not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Flag Fairway",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
