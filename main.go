package main

import (
	"assignment2/database"
	"assignment2/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
