package utils

import (
	"final-project/config"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseDriver string

const (
	Postgres DatabaseDriver = "postgres"
)

type DatabaseConnection struct {
	Driver DatabaseDriver

	Postgres *gorm.DB
}

func NewConnectionDatabase(config *config.AppConfig) *DatabaseConnection {
	var db DatabaseConnection
	db.Driver = Postgres
	db.Postgres = newPosgres(config)

	return &db
}

func newPosgres(config *config.AppConfig) *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func (db *DatabaseConnection) CloseConnection() {
	if db.Postgres != nil {
		db, _ := db.Postgres.DB()
		db.Close()
	}
}
