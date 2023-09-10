package main

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Password  string `gorm:"<-"`
	PhoneArea int
	Phone     int
	Orders    []Order
}

type Employee struct {
	gorm.Model
	WorkingHoursStart time.Time
	WorkingHoursEnd   time.Time
	BusinessID        int
	Business          Business
	Name              string
	Email             string `gorm:"uniqueIndex"`
	Password          string `gorm:"<-"`
	PhoneArea         int
	Phone             int
}

type Business struct {
	gorm.Model
	Employees   []Employee
	Name        string
	Description string
	Email       string `gorm:"uniqueIndex"`
	PhoneArea   int
	Phone       int
	Orders      []Order
}

type Order struct {
	gorm.Model
	BusinessID   uint
	Business     Business
	UserID       uint
	User         User
	Products     []OrderProduct
	isFullfilled sql.NullBool `gorm:"default:false"`
	status       int          `gorm:"default:0"` // 0: pending, 1: getting loaded, 2: on the way, 3: delivered
}

type OrderProduct struct {
	gorm.Model
	Product   Product
	ProductID int
	Order     Order
	OrderID   int
	Quantity  int
}

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Currency    string
	Quantity    int
}
