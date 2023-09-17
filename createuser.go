package main

import "github.com/gin-gonic/gin"

func CreateUser(c *gin.Context) {
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

	userdb := db.First(&User{}, User{Email: body.Email})
	employeedb := db.First(&Employee{}, Employee{User: User{Email: body.Email}})

	employeedbByBusiness := db.First(&Employee{}, Employee{BusinessID: body.BusinessID})

	// Check if user already exists, return 400 if it does
	if body.BusinessID == 0 {
		if userdb.Error == nil {
			c.Status(400)
			return
		}
	} else {
		if employeedb.Error == nil {
			c.Status(400)
			return
		}
	}

	if body.BusinessID != 0 {
		if employeedbByBusiness.Error != nil {
			var employ Employee = Employee{
				User: User{
					Name:     body.Name,
					Email:    body.Email,
					Password: body.Password,
				},
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
}
