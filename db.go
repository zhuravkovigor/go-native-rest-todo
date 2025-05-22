package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	// Загрузка .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	if user == "" || password == "" || host == "" || port == "" || name == "" {
		log.Fatal("Not all environment variables are set (DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)")
	}

	// Формирование строки подключения
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, name, sslmode)

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error while connecting to DB:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error ping to DB:", err)
	}

	createTableIfNotExists()
}

func createTableIfNotExists() {
	query := `
		DROP TABLE IF EXISTS todos;

		CREATE EXTENSION IF NOT EXISTS "pgcrypto";

		CREATE TABLE IF NOT EXISTS todos (
					id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	        title TEXT NOT NULL,
	        done BOOLEAN NOT NULL DEFAULT FALSE
	    );
		`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatal("Error while creating table:", err)
	}
}
