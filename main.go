package main

import (
	"finalProject2/database"
	"finalProject2/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
