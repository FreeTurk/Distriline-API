package routes

import (
	"distriline/account"
	"distriline/db"
	"fmt"
)

func ErrorCleaner(body any, err error) any {
	if err != nil {
		fmt.Println(err)
	}
	return body
}

func SetupGinRoutes() {
	db.Res.POST("/v1/reauthUser", account.AuthRelogin)

	db.Res.POST("/v1/getPassHash", account.GetPassHash)

	db.Res.POST("/v1/createUser", account.CreateUser)
}
