package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/watarui/go-intermediate-practice/models"
)

type HelloController struct{}

func NewHelloController() *HelloController {
	return &HelloController{}
}

func createErrorResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// HelloHandler
//
//	@Summary		GET /hello のハンドラ
//	@Description	hello world を返す
//	@Produce		json
//	@Success		200	{object}	models.HelloWorld
//	@Router			/hello [get]
func (c *HelloController) HelloHandler(w http.ResponseWriter, req *http.Request) {
	helloWorld := models.HelloWorld{
		Message: "Hello, world!",
	}

	json.NewEncoder(w).Encode(helloWorld)
}
