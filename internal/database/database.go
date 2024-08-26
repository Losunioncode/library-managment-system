package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB *sql.DB

func GetEnvDB() []string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASSWORD")
	SECRET_KEY := os.Getenv("SECRET_KEY")

	return append(make([]string, 0), DB_USER, DB_PASS, SECRET_KEY)
}

func InitDB() {
	cfg := mysql.Config{
		User:   GetEnvDB()[0],
		Passwd: GetEnvDB()[1],
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "bookshelf",
	}

	var err error

	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Successfully connected to database!")

}
