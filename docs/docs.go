// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Jintae, kim",
            "url": "http://whatissuenow.com",
            "email": "6199@outlook.kr"
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
        "/list": {
            "get": {
                "description": "현재시간 기준 3시간 전까지의 상위 이슈 단어 30개를 반환한다",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatIssueNow"
                ],
                "summary": "이슈 단어목록 API",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/win_m.WList"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/libs.APIError"
                        }
                    }
                }
            }
        },
        "/related/list/{word}": {
            "get": {
                "description": "특정 단어의 관련된 최근 트윗 목록 100개를 반환한다",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatIssueNow"
                ],
                "summary": "특정 단어의 최근 트윗 목록 API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Word",
                        "name": "word",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/win_m.RTweets"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/libs.APIError"
                        }
                    }
                }
            }
        },
        "/related/{word}": {
            "get": {
                "description": "특정 단어와 관련된 다른 단어들을 반환한다",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatIssueNow"
                ],
                "summary": "특정 단어와 연관된 단어 목록 API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Word",
                        "name": "word",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/win_m.WordsResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/libs.APIError"
                        }
                    }
                }
            }
        },
        "/stopwords": {
            "get": {
                "description": "ElasticSearch 집계시에 제외되는 불용어 목록을 반환한다",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "StopWords"
                ],
                "summary": "불용어 단어목록 API",
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
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/libs.APIError"
                        }
                    }
                }
            }
        },
        "/tag/{word}": {
            "get": {
                "description": "특정 단어의 발생지(태그) 점유율을 반환한다",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatIssueNow"
                ],
                "summary": "단어별 태그 점유율 API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Word",
                        "name": "word",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/win_m.WTag"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/libs.APIError"
                        }
                    }
                }
            }
        },
        "/wordcloud": {
            "get": {
                "description": "현재시간 기준 3시간 전까지의 상위 이슈 단어 60개를 반환한다",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatIssueNow"
                ],
                "summary": "워드 클라우드를 위한 이슈 단어목록 API",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "워드클라우드 단어 개수(최대 100개)",
                        "name": "count",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "type": "object"
                                }
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/libs.APIError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "libs.APIError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "win_m.RLinks": {
            "type": "object",
            "properties": {
                "sid": {
                    "type": "string"
                },
                "tid": {
                    "type": "string"
                }
            }
        },
        "win_m.RNodes": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "win_m.RTweets": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "win_m.WList": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "word": {
                    "type": "string"
                }
            }
        },
        "win_m.WTag": {
            "type": "object",
            "properties": {
                "percent": {
                    "type": "number"
                },
                "tag": {
                    "type": "string"
                }
            }
        },
        "win_m.WordsResponse": {
            "type": "object",
            "properties": {
                "links": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/win_m.RLinks"
                    }
                },
                "nodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/win_m.RNodes"
                    }
                },
                "rank": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "number"
                    }
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
	Version:     "1.1",
	Host:        "localhost:8000",
	BasePath:    "/api/1",
	Schemes:     []string{"http", "https"},
	Title:       "WhatIssueNow API",
	Description: "WhatIssueNow Service API",
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
	swag.Register(swag.Name, &s{})
}
