package db

import (
	models "distriline/models"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Res     = gin.Default()
	dsn     = "host=localhost user=emiroven password= dbname=distriline port=5432 sslmode=disable TimeZone=Europe/Istanbul"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
)

func SetupDB() {
	err = godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Db.AutoMigrate(&models.User{}, &models.Order{}, &models.Product{}, &models.Business{}, &models.Employee{}, &models.OrderProduct{})

}
