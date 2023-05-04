package services_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/watarui/go-intermediate-practice/services"

	_ "github.com/go-sql-driver/mysql"
)

var s *services.MyAppService

func TestMain(m *testing.M) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbConn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s = services.NewMyAppService(db)

	m.Run()
}

func BenchmarkGetArticleService(b *testing.B) {
	articleID := 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := s.GetArticleService(articleID)
		if err != nil {
			b.Error(err)
			break
		}
	}
}
