// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/load": {
            "get": {
                "security": [
                    {
                        "Authentication Api Key": []
                    }
                ],
                "description": "Load Resolvers description",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "load"
                ],
                "summary": "Load Resolvers",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/resolve/{resolver}": {
            "post": {
                "security": [
                    {
                        "Authentication Api Key": []
                    }
                ],
                "description": "Resolvers description",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resolve"
                ],
                "summary": "Execute the Resolve resolutions",
                "parameters": [
                    {
                        "type": "string",
                        "description": "resolver",
                        "name": "resolver",
                        "in": "path"
                    },
                    {
                        "description": "Parameters",
                        "name": "parameters",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.Resolve"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/resolvers/": {
            "get": {
                "security": [
                    {
                        "Authentication Api Key": []
                    }
                ],
                "description": "List Resolvers description",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resolvers"
                ],
                "summary": "List Resolvers",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "v1.Resolve": {
            "type": "object",
            "properties": {
                "context": {
                    "type": "object",
                    "additionalProperties": true
                },
                "load": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "options": {
                    "type": "object",
                    "additionalProperties": true
                },
                "resolver": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Authentication Api Key": {
            "type": "apiKey",
            "name": "X-API-Key",
            "in": "header"
        }
    },
    "x-extension-openapi": {
        "example": "value on a json format"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "FeatWS Resolver Bridge",
	Description:      "API Project that provide the data of datacenter to FeatWS",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
