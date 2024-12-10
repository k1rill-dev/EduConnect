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
        "/api/applications": {
            "post": {
                "description": "Создает новый отклик на вакансию",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "applications"
                ],
                "summary": "Создать отклик",
                "parameters": [
                    {
                        "description": "Данные отклика",
                        "name": "application",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.CreateJobApplication"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Успешно создано",
                        "schema": {
                            "$ref": "#/definitions/response.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Ошибка валидации",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/applications/{applicationId}": {
            "delete": {
                "description": "Удаляет отклик",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "applications"
                ],
                "summary": "Удалить отклик",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID отклика",
                        "name": "applicationId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешно удалено",
                        "schema": {
                            "$ref": "#/definitions/response.SuccessResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/applications/{applicationId}/status": {
            "put": {
                "description": "Обновляет статус существующего отклика",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "applications"
                ],
                "summary": "Обновить статус отклика",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID отклика",
                        "name": "applicationId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Новый статус",
                        "name": "status",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешно обновлено",
                        "schema": {
                            "$ref": "#/definitions/response.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Ошибка запроса",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/refresh-tokens": {
            "post": {
                "description": "Обновление access и refresh токенов с использованием валидного refresh токена",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Обновление Access и Refresh токенов",
                "parameters": [
                    {
                        "description": "RefreshTokens Request",
                        "name": "refreshTokens",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.RefreshTokensRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Новые access и refresh токены",
                        "schema": {
                            "$ref": "#/definitions/response.RefreshTokensResponse"
                        }
                    },
                    "400": {
                        "description": "Неверный refresh токен или ошибка валидации",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/sign-in": {
            "post": {
                "description": "Авторизация пользователя по email и паролю",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Вход пользователя",
                "parameters": [
                    {
                        "description": "Данные для входа",
                        "name": "signInRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.SignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Tokens",
                        "schema": {
                            "$ref": "#/definitions/response.SignInResponse"
                        }
                    },
                    "400": {
                        "description": "Ошибка валидации или неверные учетные данные",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/sign-out": {
            "post": {
                "description": "Завершение сессии пользователя с аннулированием токенов",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Выход пользователя",
                "parameters": [
                    {
                        "description": "Данные для выхода",
                        "name": "signOutRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.SignOutRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Пустой ответ при успешном завершении",
                        "schema": {
                            "$ref": "#/definitions/response.SignOutResponse"
                        }
                    },
                    "400": {
                        "description": "Ошибка валидации или неверные данные",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/sign-up": {
            "post": {
                "description": "Создаёт нового пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Регистрация пользователя",
                "parameters": [
                    {
                        "description": "Данные для регистрации",
                        "name": "signUpRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Tokens",
                        "schema": {
                            "$ref": "#/definitions/response.SignUpResponse"
                        }
                    },
                    "400": {
                        "description": "Ошибка валидации",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/update-user": {
            "post": {
                "description": "Обновление пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Обновления пользователя",
                "parameters": [
                    {
                        "description": "Данные для обновления",
                        "name": "signInRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Tokens",
                        "schema": {
                            "$ref": "#/definitions/response.UpdateResponse"
                        }
                    },
                    "400": {
                        "description": "Ошибка валидации или неверные учетные данные",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/jobs": {
            "post": {
                "description": "Создает новую вакансию",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Создать вакансию",
                "parameters": [
                    {
                        "description": "Данные для создания вакансии",
                        "name": "job",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.CreateJobRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Успешно создано",
                        "schema": {
                            "$ref": "#/definitions/response.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Ошибка валидации",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/jobs/filter": {
            "post": {
                "description": "Возвращает список вакансий, соответствующих заданным фильтрам",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Фильтр вакансий",
                "parameters": [
                    {
                        "description": "Фильтры вакансий",
                        "name": "filters",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/repository.JobFilters"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Номер страницы",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Количество записей на странице",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Список вакансий",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Job"
                            }
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/jobs/search": {
            "get": {
                "description": "Поиск вакансий по названию",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Поиск вакансий",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Название вакансии",
                        "name": "title",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Номер страницы",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Количество записей на странице",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Список вакансий",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Job"
                            }
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/jobs/{jobId}": {
            "get": {
                "description": "Возвращает вакансию по ее ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Получить вакансию по ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID вакансии",
                        "name": "jobId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Информация о вакансии",
                        "schema": {
                            "$ref": "#/definitions/model.Job"
                        }
                    },
                    "404": {
                        "description": "Вакансия не найдена",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Обновляет данные существующей вакансии",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Обновить вакансию",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID вакансии",
                        "name": "jobId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Данные для обновления вакансии",
                        "name": "updateJob",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UpdateJobRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешно обновлено",
                        "schema": {
                            "$ref": "#/definitions/response.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Ошибка валидации",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Job": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "employerId": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "repository.JobFilters": {
            "type": "object",
            "properties": {
                "employerId": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                }
            }
        },
        "requests.CreateJobApplication": {
            "type": "object",
            "required": [
                "companyId",
                "status",
                "studentId"
            ],
            "properties": {
                "companyId": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "studentId": {
                    "type": "string"
                }
            }
        },
        "requests.CreateJobRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "requests.RefreshTokensRequest": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "requests.SignInRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "requests.SignOutRequest": {
            "type": "object"
        },
        "requests.SignUpRequest": {
            "type": "object",
            "required": [
                "bio",
                "email",
                "firstName",
                "password",
                "picture",
                "role",
                "surname"
            ],
            "properties": {
                "bio": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "requests.UpdateJobRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "requests.UpdateRequest": {
            "type": "object",
            "required": [
                "bio",
                "email",
                "firstName",
                "picture",
                "surname"
            ],
            "properties": {
                "bio": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "response.ErrorResponse": {
            "type": "object",
            "required": [
                "error"
            ],
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "response.RefreshTokensResponse": {
            "type": "object",
            "required": [
                "access_token",
                "refresh_token"
            ],
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "response.SignInResponse": {
            "type": "object",
            "required": [
                "access_token",
                "refresh_token"
            ],
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "response.SignOutResponse": {
            "type": "object"
        },
        "response.SignUpResponse": {
            "type": "object",
            "required": [
                "access_token",
                "refresh_token"
            ],
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "response.SuccessResponse": {
            "type": "object",
            "required": [
                "message"
            ],
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "response.UpdateResponse": {
            "type": "object",
            "required": [
                "bio",
                "email",
                "firstName",
                "picture",
                "surname"
            ],
            "properties": {
                "bio": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "surname": {
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
