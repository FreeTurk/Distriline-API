package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AuthRelogin(c *gin.Context) {
	type ExpectedReq struct {
		AuthUuid string `json:"auth-uuid"`
		Email    string `json:"email"`
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

	result := map[string]interface{}{}

	userdb := db.Where("Email = ?", body.Email).First(&User{})
	employeedb := db.Where("Email = ?", body.Email).First(&Employee{})

	var options []*gorm.DB = []*gorm.DB{userdb, employeedb}

	for _, option := range options {
		if option.Error == nil {
			userdata := map[string]interface{}{}
			option.Take(&userdata)

			// check if uuid is valid
			if userdata["AuthUuid"] != body.AuthUuid {
				c.JSON(400, gin.H{
					"result": false,
				})
				return
			}

			// check if checksum is valid
			isChecksumValid, _ := CheckUserIntegrity(userdata)

			if !isChecksumValid {
				c.JSON(400, gin.H{
					"result": false,
				})
				return
			}

			// if valid, refresh the key and return the user
			option.Update("AuthUuid", uuid.New().String())
			option.Take(&result)
			c.JSON(200, gin.H{
				"result": result,
				"type":   "user",
			})
			return
		}
	}

	c.JSON(400, gin.H{
		"result": false,
	})
	return
}
