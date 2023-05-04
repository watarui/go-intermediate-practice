package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/watarui/go-intermediate-practice/api"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbHost     = os.Getenv("DB_HOST")
	dbConn     = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("failed to connect to mysql")
		return
	}

	r := api.NewRouter(db)

	log.Println("server start at port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
