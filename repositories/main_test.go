package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var testDB *sql.DB

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbHost     = os.Getenv("DB_HOST")
	dbConn     = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbDatabase)
)

func connectDB() error {
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

func setupTestData() error {
	cmd := exec.Command("mysql", "--protocol=tcp", "-h", dbHost, "-u", dbUser, dbDatabase, "--password="+dbPassword, "-e", "source ./testdata/setupDB.sql")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func cleanupDB() error {
	cmd := exec.Command("mysql", "--protocol=tcp", "-h", dbHost, "-u", dbUser, dbDatabase, "--password="+dbPassword, "-e", "source ./testdata/cleanupDB.sql")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// 全テスト共通の前処理を書く
func setup() error {
	if err := connectDB(); err != nil {
		fmt.Println("connect", err)
		return err
	}
	if err := cleanupDB(); err != nil {
		fmt.Println("cleanup", err)
		return err
	}
	if err := setupTestData(); err != nil {
		fmt.Println("setup")
		return err
	}
	return nil
}

// 前テスト共通の後処理を書く
func teardown() {
	cleanupDB()
	testDB.Close()
}

func TestMain(m *testing.M) {
	fmt.Println("setup")

	err := setup()
	if err != nil {
		os.Exit(1)
	}

	m.Run()

	teardown()
}
