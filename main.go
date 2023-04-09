package main

import (
	"tugas2/database"
	"tugas2/router"
)

func main() {
	database.StartDB()

	router.New().Run(":3000")
}