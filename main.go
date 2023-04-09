package main

import (
	"tugas2/database"
)

func main() {
	database.StartDB()

	router.New().Run(":3000")
}