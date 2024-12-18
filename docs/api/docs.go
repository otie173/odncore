// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/license/mit"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/player/auth": {
            "post": {
                "description": "Register new player or login existing player",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "player"
                ],
                "summary": "Authenticate player",
                "parameters": [
                    {
                        "description": "Player authentication data",
                        "name": "playerAuth",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.PlayerAuth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "FAIL",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/player/getpdata": {
            "get": {
                "description": "Get player's saved data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "player"
                ],
                "summary": "Get player data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Player nickname",
                        "name": "Session-Nickname",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    },
                    "404": {
                        "description": "Player not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/player/loadpdata": {
            "post": {
                "description": "Save player's current state",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "player"
                ],
                "summary": "Save player data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Player nickname",
                        "name": "Session-Nickname",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Player data to save",
                        "name": "playerData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/server/info": {
            "get": {
                "description": "Get information about the server",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "server"
                ],
                "summary": "Get server info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.ServerInfo"
                        }
                    }
                }
            }
        },
        "/server/status": {
            "get": {
                "description": "Get current server status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "server"
                ],
                "summary": "Get server status",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.ServerStatus"
                        }
                    }
                }
            }
        },
        "/world/getworld": {
            "get": {
                "description": "Get world in binary format",
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "world"
                ],
                "summary": "Get world",
                "responses": {
                    "200": {
                        "description": "World binary data",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    }
                }
            }
        },
        "/world/getworldinfo": {
            "get": {
                "description": "Get world info in json format",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "world"
                ],
                "summary": "Get world info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/world.WorldInfo"
                        }
                    }
                }
            }
        },
        "/world/loadid": {
            "post": {
                "description": "Load world id data",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "world"
                ],
                "summary": "Load world ID",
                "parameters": [
                    {
                        "description": "World ID data",
                        "name": "idData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/world/loadworld": {
            "post": {
                "description": "Load world from binary file",
                "consumes": [
                    "application/octet-stream"
                ],
                "tags": [
                    "world"
                ],
                "summary": "Load world",
                "parameters": [
                    {
                        "format": "binary",
                        "description": "World binary data",
                        "name": "worldData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/world/loadworldinfo": {
            "post": {
                "description": "Load world info in json format",
                "consumes": [
                    "application/octet-stream"
                ],
                "tags": [
                    "world"
                ],
                "summary": "Load world info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.PlayerAuth": {
            "type": "object",
            "properties": {
                "nickname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "server.ServerInfo": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "max_players": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "players_connected": {
                    "type": "integer"
                }
            }
        },
        "server.ServerStatus": {
            "type": "object",
            "properties": {
                "id_waiting": {
                    "type": "boolean"
                },
                "world_waiting": {
                    "type": "boolean"
                }
            }
        },
        "world.WorldInfo": {
            "type": "object",
            "properties": {
                "axes_count": {
                    "type": "integer"
                },
                "big_stones_count": {
                    "type": "integer"
                },
                "bones_generated": {
                    "type": "boolean"
                },
                "pickaxes_count": {
                    "type": "integer"
                },
                "saplings_count": {
                    "type": "integer"
                },
                "seeds_count": {
                    "type": "integer"
                },
                "shovels_count": {
                    "type": "integer"
                },
                "small_stones_count": {
                    "type": "integer"
                },
                "structures_generated": {
                    "type": "boolean"
                },
                "trees_count": {
                    "type": "integer"
                }
            }
        }
    },
    "tags": [
        {
            "description": "Server management endpoints",
            "name": "server"
        },
        {
            "description": "World management endpoints",
            "name": "world"
        },
        {
            "description": "Player management endpoints",
            "name": "player"
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "0.0.0.0:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Odncore API",
	Description:      "API server for Odinbit game",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}