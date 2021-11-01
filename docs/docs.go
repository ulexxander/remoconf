// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/projects": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "operationId": "GetProjectsAll",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/storage.Project"
                            }
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/restapi.ResponseError"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "operationId": "PostProject",
                "parameters": [
                    {
                        "description": "Project Create Params",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/storage.ProjectCreateParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/storage.CreatedItem"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/restapi.ResponseError"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "operationId": "PostUser",
                "parameters": [
                    {
                        "description": "User Create Params",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/storage.UserCreateParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/storage.CreatedItem"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/restapi.ResponseError"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "operationId": "GetUserByID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/storage.User"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/restapi.ResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "restapi.ResponseError": {
            "type": "object",
            "properties": {
                "Error": {
                    "type": "string"
                }
            }
        },
        "storage.CreatedItem": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                }
            }
        },
        "storage.Project": {
            "type": "object",
            "properties": {
                "CreatedAt": {
                    "type": "string"
                },
                "CreatedBy": {
                    "type": "integer"
                },
                "Description": {
                    "type": "string"
                },
                "ID": {
                    "type": "integer"
                },
                "Title": {
                    "type": "string"
                },
                "UpdatedAt": {
                    "type": "string"
                },
                "UpdatedBy": {
                    "type": "integer"
                }
            }
        },
        "storage.ProjectCreateParams": {
            "type": "object",
            "properties": {
                "CreatedBy": {
                    "type": "integer"
                },
                "Description": {
                    "type": "string"
                },
                "Title": {
                    "type": "string"
                }
            }
        },
        "storage.User": {
            "type": "object",
            "properties": {
                "CreatedAt": {
                    "type": "string"
                },
                "ID": {
                    "type": "integer"
                },
                "Login": {
                    "type": "string"
                },
                "Password": {
                    "type": "string"
                },
                "UpdatedAt": {
                    "type": "string"
                }
            }
        },
        "storage.UserCreateParams": {
            "type": "object",
            "properties": {
                "Login": {
                    "type": "string"
                },
                "Password": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
