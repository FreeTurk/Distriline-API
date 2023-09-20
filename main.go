package main

import (
	"distriline/db"
	"distriline/routes"
)

func main() {
	db.SetupDB()

	routes.SetupGinRoutes()
	db.Res.Run(":6151")
}
