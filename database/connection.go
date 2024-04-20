package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DBSTRING := os.Getenv("DB_STRING")


	db, err := sql.Open("sqlite3", DBSTRING)

	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
