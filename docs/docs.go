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
        "/Reservation/delete/{id}": {
            "delete": {
                "description": "Delete page",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reservation"
                ],
                "summary": "Delete Reservation",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Reservation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Delete Successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Error while Deleted",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/Reservation/getall": {
            "get": {
                "description": "GetAll page",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reservation"
                ],
                "summary": "GetAll Reservation",
                "responses": {
                    "200": {
                        "description": "GetAll Successful",
                        "schema": {
                            "$ref": "#/definitions/genprotos.GetAllReservations"
                        }
                    },
                    "401": {
                        "description": "Error while GetAlld",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/Reservation/getbyid/{id}": {
            "get": {
                "description": "GetById page",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reservation"
                ],
                "summary": "GetById Reservation",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Reservation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetById Successful",
                        "schema": {
                            "$ref": "#/definitions/genprotos.Reservation"
                        }
                    },
                    "401": {
                        "description": "Error while GetByIdd",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/Reservation/update/{id}": {
            "put": {
                "description": "Update page",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reservation"
                ],
                "summary": "Update Reservation",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Reservation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update",
                        "name": "Update",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.Reservation"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Update Successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Error while created",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/reservation/create/{user_id}/{restaurant_id}": {
            "post": {
                "description": "Create page",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reservation"
                ],
                "summary": "Create Reservation",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Restoran ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User id",
                        "name": "restaurant_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Create",
                        "name": "Create",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.Reservation"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Create Successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Error while Created",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "genprotos.GetAllReservations": {
            "type": "object",
            "properties": {
                "reservations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/genprotos.Reservation"
                    }
                }
            }
        },
        "genprotos.Reservation": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "reservation_time": {
                    "type": "string"
                },
                "restaurant_id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
