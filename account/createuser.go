package account

import (
	"distriline/db"
	"distriline/models"

	"github.com/gin-gonic/gin"
)

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

	userdb := db.Db.First(&models.User{}, models.User{Email: body.Email})
	employeedb := db.Db.First(&models.Employee{}, models.Employee{User: models.User{Email: body.Email}})

	employeedbByBusiness := db.Db.First(&models.Employee{}, models.Employee{BusinessID: body.BusinessID})

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
			var employ models.Employee = models.Employee{
				User: models.User{
					Name:     body.Name,
					Email:    body.Email,
					Password: body.Password,
				},
				BusinessID: body.BusinessID,
			}
			db.Db.Create(&employ)

		}
	} else {
		var user models.User = models.User{
			Name:     body.Name,
			Email:    body.Email,
			Password: body.Password,
		}
		db.Db.Create(&user)
	}
	c.Status(200)
}
