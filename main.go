package main

import (
	"Latihan1/config"
	"Latihan1/routers"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Dbname)

	var err error
	config.Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer config.Db.Close()

	err = config.Db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database successfully!")
	routers.BookRouter()
}
