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
        "/read/num/exponentialrate/{name}": {
            "get": {
                "description": "Read rate estimates of a specific statistic. Not suitable for accuracy, but great for quickly comparing stats.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Read all rate estimates for a stat name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Stat name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "/read/num/exponentialrate/{name}/{label}": {
            "get": {
                "description": "Read rate estimate of a specific statistic. Not suitable for accuracy, but great for quickly comparing stats.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Read an exponential *estimate* of a numeric rate of change",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Stat name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Stat label",
                        "name": "label",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "/read/num/value/{name}": {
            "get": {
                "description": "Read a specific statistic",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Read a specific statistic",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Stat name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/api.JSONNumReadResult"
                        }
                    }
                }
            }
        },
        "/read/num/value/{name}/{label}": {
            "get": {
                "description": "Read a specific statistic",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Read a specific statistic",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Stat name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Stat label",
                        "name": "label",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "/register/num/decrease": {
            "post": {
                "description": "Decrease a numeric stat",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Decrease a numeric stat",
                "parameters": [
                    {
                        "description": "Stat and value",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.NumStatRegistration"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "/register/num/increase": {
            "post": {
                "description": "Increase a numeric stat",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Increase a numeric stat",
                "parameters": [
                    {
                        "description": "Stat and value",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.NumStatRegistration"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "/register/num/set": {
            "post": {
                "description": "Set a numeric stat",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Set numeric stat",
                "parameters": [
                    {
                        "description": "Stat and value",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.NumStatRegistration"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.JSONNumReadResult": {
            "type": "object",
            "properties": {
                "stat_name1": {
                    "type": "number"
                },
                "stat_name2": {
                    "type": "number"
                }
            }
        },
        "api.NumStatRegistration": {
            "type": "object",
            "properties": {
                "label": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
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
