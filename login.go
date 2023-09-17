package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetPassHash(c *gin.Context) {
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

	userdb := db.First(&User{}, User{Email: body.Email})
	employeedb := db.First(&Employee{}, Employee{User: User{Email: body.Email}})

	var options []*gorm.DB = []*gorm.DB{userdb, employeedb}

	// check if user exists
	if userdb.Error != nil && employeedb.Error != nil {
		fmt.Println("user not found")
		c.Status(400)
		return
	}

	result := map[string]interface{}{}

	for _, option := range options {
		if option.Error == nil {
			// update auth uuid
			userdb.Update("AuthUuid", uuid.New().String())

			// get user
			userdb.Take(&result)

			// refresh checksums
			_, checksum := CheckUserIntegrity(result)
			userdb.Update("Checksum", checksum)

			// reget user
			userdb.Take(&result)

			var usertype string

			if option == userdb {
				usertype = "user"
			} else {
				usertype = "employee"
			}

			// return user
			c.JSON(200, gin.H{
				"result": result,
				"type":   usertype,
			})
			return
		}
	}

	return
}
