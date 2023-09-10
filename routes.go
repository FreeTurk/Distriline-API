package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorCleaner(body any, err error) any {
	if err != nil {
		fmt.Println(err)
	}
	return body
}

func setupGinRoutes() {
	res.POST("/v1/getPassHash", func(c *gin.Context) {
		type ExpectedReq struct {
			Email string `json:"email"`
		}

		var body ExpectedReq

		err := c.BindJSON(&body)
		if err != nil {
			c.JSON(404, gin.H{
				"code":    404,
				"message": "bad request",
			})
			return
		}

		fmt.Println(body)
		// check if user exists
		if (db.Where("Email = ?", body.Email).First(&User{}).Error != nil && db.Where("Email = ?", body.Email).First(&Employee{}).Error != nil) {
			fmt.Println("user not found")
			c.Status(400)
			return
		}

		result := map[string]interface{}{}
		if (db.Where("Email = ?", body.Email).First(&User{}).Error == nil) {
			db.Where("Email = ?", body.Email).Model(&User{}).Take(&result)
			c.JSON(200, gin.H{
				"result": result,
				"type":   "user",
			})
		} else if (db.Where("Email = ?", body.Email).First(&Employee{}).Error == nil) {
			db.Where("Email = ?", body.Email).Model(&Employee{}).Take(&result)
			c.JSON(200, gin.H{
				"result": result,
				"type":   "employee",
			})
		}
		return
	})

	res.POST("/v1/createUser", func(c *gin.Context) {
		type ExpectedReq struct {
			Name       string `json:"name"`
			Email      string `json:"email"`
			Password   string `json:"password"`
			BusinessID int    `json:"businessID"`
		}

		var body ExpectedReq

		err := c.BindJSON(&body)

		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"message": "bad request",
			})
			return
		}

		// Check if user already exists, return 400 if it does
		if body.BusinessID == 0 {
			if (db.Where("Email = ?", body.Email).First(&User{}).Error == nil) {
				c.Status(400)
				return
			}
		} else {
			if (db.Where("Email = ?", body.Email).First(&Employee{}).Error == nil) {
				c.Status(400)
				return
			}
		}

		if body.BusinessID != 0 {
			if (db.Where("RegistrationID = ?", body.BusinessID).First(&Business{}).Error != nil) {
				var employ Employee = Employee{
					Name:       body.Name,
					Email:      body.Email,
					Password:   body.Password,
					BusinessID: body.BusinessID,
				}
				db.Create(&employ)
			}
		} else {
			var user User = User{
				Name:     body.Name,
				Email:    body.Email,
				Password: body.Password,
			}
			db.Create(&user)
		}
		c.Status(200)
	})
}
