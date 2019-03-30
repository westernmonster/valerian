// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-03-30 23:59:08.062385 +0800 CST m=+0.031064366

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "飞行百科 API",
        "contact": {},
        "license": {},
        "version": "1.0"
    },
    "host": "www.flywk.com",
    "basePath": "/api/v1",
    "paths": {
        "/accounts": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "注册请求",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.RegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功"
                    },
                    "400": {
                        "description": "验证失败"
                    },
                    "500": {
                        "description": "服务器端错误"
                    }
                }
            }
        },
        "/accounts/attributes": {
            "put": {
                "description": "更改用户资料，需要登录验证 （JWT Token 或 Cookie）",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "更改用户资料",
                "parameters": [
                    {
                        "description": "请求",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.UpdateProfileReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功"
                    },
                    "400": {
                        "description": "验证失败"
                    },
                    "500": {
                        "description": "服务器端错误"
                    }
                }
            }
        },
        "/accounts/attributes/change_password": {
            "put": {
                "description": "更改密码，API会验证登录 （JWT Token 或 Cookie）",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "更改密码",
                "parameters": [
                    {
                        "description": "请求",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ChangePasswordReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功"
                    },
                    "400": {
                        "description": "验证失败"
                    },
                    "500": {
                        "description": "服务器端错误"
                    }
                }
            }
        },
        "/accounts/attributes/forget_password": {
            "put": {
                "description": "忘记密码，此为重设密码第一步，提交用户标识（手机号、邮箱），和用户输入的验证码进行验证，并返回一个 Session ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "忘记密码",
                "parameters": [
                    {
                        "description": "请求",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ForgetPasswordReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "SessionID"
                    },
                    "400": {
                        "description": "验证失败"
                    },
                    "500": {
                        "description": "服务器端错误"
                    }
                }
            }
        },
        "/accounts/attributes/reset_password": {
            "put": {
                "description": "重设密码第二步，传入新密码和Session ID，如果返回的Code值为307，则表示Session已经失效，前端可以根据这个值做对应的处理",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "重设密码",
                "parameters": [
                    {
                        "description": "请求",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ForgetPasswordReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功，如果返回的Code值为307，则表示Session已经失效，前端可以根据这个值做对应的处理"
                    },
                    "400": {
                        "description": "验证失败"
                    },
                    "500": {
                        "description": "服务器端错误"
                    }
                }
            }
        },
        "/session": {
            "post": {
                "description": "用户登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "session"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户登录",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.LoginResult"
                        }
                    },
                    "400": {
                        "description": "验证失败"
                    },
                    "500": {
                        "description": "服务器端错误"
                    }
                }
            }
        },
        "/valcodes": {
            "post": {
                "description": "请求验证码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "common"
                ],
                "summary": "请求验证码",
                "parameters": [
                    {
                        "description": "请求",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.RequestValcodeReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功"
                    },
                    "400": {
                        "description": "验证失败"
                    },
                    "500": {
                        "description": "服务器端错误"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ChangePasswordReq": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                }
            }
        },
        "models.ForgetPasswordReq": {
            "type": "object",
            "properties": {
                "identity": {
                    "type": "string"
                },
                "valcode": {
                    "type": "string"
                }
            }
        },
        "models.LoginReq": {
            "type": "object",
            "properties": {
                "identity": {
                    "description": "Identity 登录标识，可以传入邮件或手机号，请在提交前进行验证",
                    "type": "string"
                },
                "password": {
                    "description": "Password 密码，服务端不保存密码的明文值，请在提交前进行 MD5 哈希",
                    "type": "string"
                },
                "source": {
                    "description": "Source 来源，1:Web, 2:iOS; 3:Android",
                    "type": "integer"
                }
            }
        },
        "models.LoginResult": {
            "type": "object",
            "properties": {
                "role": {
                    "description": "用户角色，用于客户端权限管理",
                    "type": "string"
                },
                "token": {
                    "description": "Token JWT Token， 请在 HTTP 请求头中添加\n例子： Authorization: Bearer  TJVA95OrM7E20RMHrHDcEfxjoYZgeFONFh7HgQ",
                    "type": "string"
                }
            }
        },
        "models.RegisterReq": {
            "type": "object",
            "properties": {
                "identity": {
                    "description": "用户标识， 可以为邮件或手机号码",
                    "type": "string"
                },
                "password": {
                    "description": "密码 后端不保存明文密码，请于前端求得当前密码MD5哈希值后发送给后端",
                    "type": "string"
                },
                "source": {
                    "description": "Source 来源，1:Web, 2:iOS; 3:Android",
                    "type": "integer"
                },
                "valcode": {
                    "description": "验证码 6位数字",
                    "type": "string"
                }
            }
        },
        "models.RequestValcodeReq": {
            "type": "object",
            "properties": {
                "code_type": {
                    "description": "验证码类型, 1为注册验证码, 2为重置密码验证码",
                    "type": "integer"
                },
                "identity": {
                    "description": "用户标识, 可以为邮件或手机号码",
                    "type": "string"
                }
            }
        },
        "models.UpdateProfileReq": {
            "type": "object",
            "properties": {
                "avatar": {
                    "description": "用户头像\n如果无需更改该字段，在提交JSON请求中请不要包含该字段",
                    "type": "string"
                },
                "birth_day": {
                    "description": "出生日",
                    "type": "integer"
                },
                "birth_month": {
                    "description": "出生月",
                    "type": "integer"
                },
                "birth_year": {
                    "description": "出生年\n如果无需更改该字段，在提交JSON请求中请不要包含该字段",
                    "type": "integer"
                },
                "gener": {
                    "description": "用户性别， 1 为男，2 为女\n如果无需更改该字段，在提交JSON请求中请不要包含该字段",
                    "type": "integer"
                },
                "introduction": {
                    "description": "个性签名\n如果无需更改该字段，在提交JSON请求中请不要包含该字段",
                    "type": "string"
                },
                "location": {
                    "description": "地区\n如果无需更改该字段，在提交JSON请求中请不要包含该字段",
                    "type": "integer"
                },
                "password": {
                    "description": "密码\n如果无需更改该字段，在提交JSON请求中请不要包含该字段",
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
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
