// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "Lists all recipes, paginated",
                "tags": [
                    "recipes"
                ],
                "summary": "Lists all recipes",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Number of items per page",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Recipes",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.RecipeInfo"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid pagination parameters",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error while processing request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Adds a new recipe",
                "tags": [
                    "recipes"
                ],
                "summary": "Adds a new recipe",
                "parameters": [
                    {
                        "description": "New recipe data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.NewRecipe"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "New recipe ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid data format",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error while processing request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/{id}": {
            "get": {
                "description": "Gets a recipe by ID",
                "tags": [
                    "recipes"
                ],
                "summary": "Gets a recipe by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Recipe ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Recipe"
                        }
                    },
                    "400": {
                        "description": "Recipe ID is required",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error while processing request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Updates a recipe",
                "tags": [
                    "recipes"
                ],
                "summary": "Updates a recipe",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Recipe ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Recipe data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.NewRecipeDataNoID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Recipe updated",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid data format",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error while processing request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a recipe by ID",
                "tags": [
                    "recipes"
                ],
                "summary": "Deletes a recipe by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Recipe ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Recipe deleted",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Recipe ID is required",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error while processing request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Ingredient": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "quantity": {
                    "type": "number"
                },
                "unit": {
                    "type": "string"
                }
            }
        },
        "models.NewRecipe": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "cook_time": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Ingredient"
                    }
                },
                "instructions": {
                    "type": "string"
                },
                "prep_time": {
                    "type": "integer"
                },
                "servings": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.NewRecipeDataNoID": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "cook_time": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Ingredient"
                    }
                },
                "instructions": {
                    "type": "string"
                },
                "prep_time": {
                    "type": "integer"
                },
                "servings": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.Recipe": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "cook_time": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Ingredient"
                    }
                },
                "instructions": {
                    "type": "string"
                },
                "prep_time": {
                    "type": "integer"
                },
                "servings": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.RecipeInfo": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/recipes",
	Schemes:          []string{"http"},
	Title:            "Recipe Management System",
	Description:      "API for Recipe Management System",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
