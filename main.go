package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	res     = gin.Default()
	dsn     = "host=localhost user=emiroven password=1903bjksk dbname=distriline port=5432 sslmode=disable TimeZone=Europe/Istanbul"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
)

func main() {
	db.AutoMigrate(&User{}, &Order{}, &Product{}, &Business{}, &Employee{})

	fmt.Println("done.")

	setupGinRoutes()

	res.Run(":6151")
}
