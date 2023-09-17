package main

import (
	"fmt"
)

func ErrorCleaner(body any, err error) any {
	if err != nil {
		fmt.Println(err)
	}
	return body
}

func setupGinRoutes() {
	res.POST("/v1/reauthUser", AuthRelogin)

	res.POST("/v1/getPassHash", GetPassHash)

	res.POST("/v1/createUser", CreateUser)
}
