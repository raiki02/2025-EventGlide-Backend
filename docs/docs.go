// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/act/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Activity"
                ],
                "summary": "创建活动",
                "parameters": [
                    {
                        "description": "活动",
                        "name": "activity",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Activity"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/act/date/{date}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Activity"
                ],
                "summary": "通过日期查找活动",
                "parameters": [
                    {
                        "type": "string",
                        "description": "日期",
                        "name": "date",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/act/draft": {
            "post": {
                "description": "not finished",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Activity"
                ],
                "summary": "创建活动草稿",
                "parameters": [
                    {
                        "description": "活动草稿",
                        "name": "draft",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ActivityDraft"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/act/foreign": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Activity"
                ],
                "summary": "通过是否为外部活动查找活动",
                "parameters": [
                    {
                        "type": "string",
                        "description": "类型",
                        "name": "type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/act/host": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Activity"
                ],
                "summary": "通过主办方查找活动",
                "parameters": [
                    {
                        "type": "string",
                        "description": "主办方",
                        "name": "host",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/act/location": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Activity"
                ],
                "summary": "通过地点查找活动",
                "parameters": [
                    {
                        "type": "string",
                        "description": "地点",
                        "name": "location",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/act/name": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Activity"
                ],
                "summary": "通过名称查找活动",
                "parameters": [
                    {
                        "type": "string",
                        "description": "名称",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/act/signup": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Activity"
                ],
                "summary": "通过是否需要报名查找活动",
                "parameters": [
                    {
                        "type": "string",
                        "description": "类型",
                        "name": "type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/act/time": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Activity"
                ],
                "summary": "通过时间查找活动",
                "parameters": [
                    {
                        "type": "string",
                        "description": "开始时间",
                        "name": "start_time",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "结束时间",
                        "name": "end_time",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/act/type": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Activity"
                ],
                "summary": "通过类型查找活动",
                "parameters": [
                    {
                        "type": "string",
                        "description": "类型",
                        "name": "type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/comment/answer": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "回复评论",
                "parameters": [
                    {
                        "description": "评论",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Comment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/comment/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "创建评论",
                "parameters": [
                    {
                        "description": "评论",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Comment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/comment/delete": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "删除评论",
                "parameters": [
                    {
                        "type": "string",
                        "description": "评论ID",
                        "name": "comment_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/number/comments": {
            "post": {
                "description": "not finished",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Number"
                ],
                "summary": "评论数控制",
                "parameters": [
                    {
                        "description": "评论数",
                        "name": "comments_num",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.NumberReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/number/likes": {
            "post": {
                "description": "not finished",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Number"
                ],
                "summary": "点赞数控制",
                "parameters": [
                    {
                        "description": "点赞数",
                        "name": "likes_num",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.NumberReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/post/all": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "获取所有帖子",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/post/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "创建帖子",
                "parameters": [
                    {
                        "description": "帖子",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Post"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/post/delete": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "删除帖子",
                "parameters": [
                    {
                        "description": "帖子",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Post"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/post/find": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "通过帖子名查找帖子",
                "parameters": [
                    {
                        "type": "string",
                        "description": "帖子名",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/user/avatar": {
            "post": {
                "description": "not finished",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "更新头像",
                "parameters": [
                    {
                        "type": "string",
                        "description": "学号",
                        "name": "sid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "图片0",
                        "name": "file0",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/user/info": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "学号",
                        "name": "sid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "学号",
                        "name": "studentid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/user/logout": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "登出",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/user/search/act": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "搜索用户活动",
                "parameters": [
                    {
                        "type": "string",
                        "description": "学号",
                        "name": "sid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "关键字",
                        "name": "keyword",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/user/search/post": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "搜索用户帖子",
                "parameters": [
                    {
                        "type": "string",
                        "description": "学号",
                        "name": "sid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "关键字",
                        "name": "keyword",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        },
        "/user/username": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "更新用户名",
                "parameters": [
                    {
                        "type": "string",
                        "description": "学号",
                        "name": "sid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "新用户名",
                        "name": "newname",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Resp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Activity": {
            "type": "object",
            "properties": {
                "audition": {
                    "type": "string"
                },
                "bid": {
                    "type": "string"
                },
                "capacity": {
                    "description": "descriptive",
                    "type": "integer"
                },
                "comments": {
                    "type": "integer"
                },
                "created_at": {
                    "description": "活动id",
                    "type": "string"
                },
                "creator_id": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "details": {
                    "type": "string"
                },
                "end_time": {
                    "type": "string"
                },
                "host": {
                    "type": "string"
                },
                "identification": {
                    "description": "audit",
                    "type": "string"
                },
                "if_register": {
                    "type": "string"
                },
                "images": {
                    "type": "string"
                },
                "likes": {
                    "description": "interactive",
                    "type": "integer"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "register_method": {
                    "type": "string"
                },
                "start_time": {
                    "description": "complex",
                    "type": "string"
                },
                "type": {
                    "description": "divided by function\nbasic",
                    "type": "string"
                }
            }
        },
        "model.ActivityDraft": {
            "type": "object",
            "properties": {
                "capacity": {
                    "description": "descriptive",
                    "type": "integer"
                },
                "created_at": {
                    "description": "活动id",
                    "type": "string"
                },
                "creator_id": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "details": {
                    "type": "string"
                },
                "end_time": {
                    "type": "string"
                },
                "host": {
                    "type": "string"
                },
                "if_register": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "register_method": {
                    "type": "string"
                },
                "start_time": {
                    "description": "complex",
                    "type": "string"
                },
                "type": {
                    "description": "divided by function\nbasic",
                    "type": "string"
                }
            }
        },
        "model.Comment": {
            "type": "object",
            "properties": {
                "answers": {
                    "type": "integer"
                },
                "comment_id": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "likes": {
                    "type": "integer"
                },
                "parent_id": {
                    "type": "string"
                },
                "poster_id": {
                    "type": "string"
                },
                "target_id": {
                    "type": "string"
                }
            }
        },
        "model.Post": {
            "type": "object",
            "properties": {
                "bid": {
                    "type": "integer"
                },
                "comments": {
                    "type": "integer"
                },
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "img_urls": {
                    "type": "string"
                },
                "likes": {
                    "type": "integer"
                },
                "poster_id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "req.NumberReq": {
            "type": "object",
            "properties": {
                "excuter_id": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                },
                "topic": {
                    "type": "string"
                }
            }
        },
        "resp.Resp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
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
	Title:            "EventGlide API",
	Description:      "校灵通 API 文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
