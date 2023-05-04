package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/watarui/go-intermediate-practice/api/middlewares"
	"github.com/watarui/go-intermediate-practice/controllers"
	"github.com/watarui/go-intermediate-practice/services"
)

func NewRouter(db *sql.DB) *mux.Router {
	s := services.NewMyAppService(db)
	ac := controllers.NewArticleController(s)
	cc := controllers.NewCommentController(s)
	hc := controllers.NewHelloController()

	r := mux.NewRouter()

	r.HandleFunc("/hello", hc.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", ac.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", ac.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", ac.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", ac.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", cc.PostCommentHandler).Methods(http.MethodPost)

	r.PathPrefix("/docs/").Handler(http.StripPrefix("/docs", http.FileServer(http.Dir("docs/"))))
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8081/docs/swagger.yaml"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("full"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	r.Use(middlewares.LoggingMiddleware)
	r.Use(middlewares.AuthMiddleware)

	return r
}
