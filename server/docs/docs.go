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
        "/login/github": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "LoginRequest"
                ],
                "summary": "GitHub Login",
                "operationId": "LoginGitHub",
                "parameters": [
                    {
                        "description": "LoginGitHubRequest",
                        "name": "q",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.LoginGitHubRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.LoginGitHubResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/rooms": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Room"
                ],
                "summary": "List Room",
                "operationId": "ListRoom",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ListRoomResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
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
                "tags": [
                    "Room"
                ],
                "summary": "Create Room",
                "operationId": "CreateRoom",
                "parameters": [
                    {
                        "description": "create room request",
                        "name": "b",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CreateRoomRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.CreateRoomResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/rooms/{roomId}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Room"
                ],
                "summary": "Get Room",
                "operationId": "GetRoom",
                "parameters": [
                    {
                        "type": "string",
                        "description": "room ID",
                        "name": "roomId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.GetRoomResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Room"
                ],
                "summary": "Update Room",
                "operationId": "UpdateRoom",
                "parameters": [
                    {
                        "type": "string",
                        "description": "roomID path param",
                        "name": "roomId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "create room request",
                        "name": "b",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.UpdateRoomRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.CreateRoomResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Room"
                ],
                "summary": "Delete Room",
                "operationId": "DeleteRoom",
                "parameters": [
                    {
                        "type": "string",
                        "description": "roomID path param",
                        "name": "roomId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.DeleteRoomResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/rooms/{roomId}/members": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Room"
                ],
                "summary": "Join Room",
                "operationId": "JoinRoom",
                "parameters": [
                    {
                        "type": "string",
                        "description": "roomID path param",
                        "name": "roomId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.JoinRoomResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Room"
                ],
                "summary": "Leave Room",
                "operationId": "LeaveRoom",
                "parameters": [
                    {
                        "type": "string",
                        "description": "roomID path param",
                        "name": "roomId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.LeaveRoomResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/skilltags": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SkillTag"
                ],
                "summary": "Serch SkillTag",
                "operationId": "SerchSkillTag",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "tag",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.SearchSkillTagResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/users/{userId}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get User",
                "operationId": "GetUser",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.GetUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update User",
                "operationId": "UpdateUser",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update user request",
                        "name": "b",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.UpdateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.CreateRoomRequest": {
            "type": "object",
            "properties": {
                "aimSkills": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "createdBy": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "controller.CreateRoomResponse": {
            "type": "object",
            "properties": {
                "aimTags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Skill"
                    }
                },
                "description": {
                    "type": "string"
                },
                "members": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.DisplayUser"
                    }
                },
                "ownerId": {
                    "type": "string"
                },
                "roomId": {
                    "type": "string"
                }
            }
        },
        "controller.DeleteRoomResponse": {
            "type": "object",
            "properties": {
                "string": {
                    "type": "string"
                }
            }
        },
        "controller.GetRoomResponse": {
            "type": "object",
            "properties": {
                "aimTags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Skill"
                    }
                },
                "description": {
                    "type": "string"
                },
                "members": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.DisplayUser"
                    }
                },
                "ownerId": {
                    "type": "string"
                },
                "roomId": {
                    "type": "string"
                }
            }
        },
        "controller.GetUserResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "icon": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "usedSkills": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Skill"
                    }
                },
                "wantLeanSkills": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Skill"
                    }
                }
            }
        },
        "controller.JoinRoomResponse": {
            "type": "object",
            "properties": {
                "aimTags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Skill"
                    }
                },
                "description": {
                    "type": "string"
                },
                "members": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.DisplayUser"
                    }
                },
                "ownerId": {
                    "type": "string"
                },
                "roomId": {
                    "type": "string"
                }
            }
        },
        "controller.LeaveRoomResponse": {
            "type": "object",
            "properties": {
                "string": {
                    "type": "string"
                }
            }
        },
        "controller.ListRoomResponse": {
            "type": "object",
            "properties": {
                "rooms": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controller.GetRoomResponse"
                    }
                }
            }
        },
        "controller.LoginGitHubRequest": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                }
            }
        },
        "controller.LoginGitHubResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "icon": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "controller.SearchSkillTagResponse": {
            "type": "object",
            "properties": {
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "controller.UpdateRoomRequest": {
            "type": "object",
            "properties": {
                "aimSkills": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "roomID": {
                    "type": "string"
                }
            }
        },
        "controller.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "usedSkills": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "userID": {
                    "type": "string"
                },
                "wantLearnSkills": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "controller.UpdateUserResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "icon": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "usedSkills": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Skill"
                    }
                },
                "wantLeanSkills": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Skill"
                    }
                }
            }
        },
        "echo.HTTPError": {
            "type": "object",
            "properties": {
                "message": {}
            }
        },
        "entity.DisplayUser": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "icon": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entity.Skill": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "submarine-api",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
