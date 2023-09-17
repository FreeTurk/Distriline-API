package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	res     = gin.Default()
	dsn     = "host=localhost user=emiroven password=1903bjksk dbname=distriline port=5432 sslmode=disable TimeZone=Europe/Istanbul"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
)

func main() {

	// Load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// All the models are defined in models.go
	// Given here to push to the database
	db.AutoMigrate(&User{}, &Order{}, &Product{}, &Business{}, &Employee{}, &OrderProduct{})

	fmt.Println("done.")

	setupGinRoutes()

	res.Run(":6151")
}
