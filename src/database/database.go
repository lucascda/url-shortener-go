package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", "localhost", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Print("connected to database")
	if err != nil {
		log.Fatal("Error connecting to database.")
	}

	DB = db
	if DB == nil {
		panic("falha ao iniciar variavel db")
	}

}

func GetDB() *gorm.DB {
	return DB
}
