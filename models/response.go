package models

type HelloWorld struct {
	Message string `json:"message" example:"Hello, world!"`
}

type InternalServerError struct {
	Message string `json:"message" example:"fail to encode json"`
}
