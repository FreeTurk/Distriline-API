package models

import (
	"database/sql"
	"time"

	gorm "gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Password  string `gorm:"<-"`
	AuthUuid  string
	PhoneArea int
	Phone     int
	Orders    []Order
	Checksum  string
}

type Employee struct {
	gorm.Model
	User              User `gorm:"embedded"`
	WorkingHoursStart time.Time
	WorkingHoursEnd   time.Time
	BusinessID        int
}

type Business struct {
	gorm.Model
	Employees   []Employee
	Name        string
	Description string
	Email       string `gorm:"uniqueIndex"`
	PhoneArea   int
	Phone       int
	Products    []Product
	Orders      []Order
}

type Order struct {
	gorm.Model
	BusinessID   uint
	UserID       uint
	User         User
	Products     []OrderProduct
	isFullfilled sql.NullBool `gorm:"default:false"`
	status       int          `gorm:"default:0"` // 0: pending, 1: getting loaded, 2: on the way, 3: delivered
}

type OrderProduct struct {
	gorm.Model
	Quantity  int
	ProductID int
	OrderID   int
}

type Product struct {
	gorm.Model
	Name        string
	BusinessID  int
	Description string
	Price       float64
	Currency    string
	Unit        string
	Interval    int
}
