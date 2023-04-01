package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"sesi_8/config"
	"sesi_8/model"
	"sesi_8/repository"
	"sesi_8/route"
	"sesi_8/service"
)

var router = gin.New()

func StartApplication() {
	db := config.PSQL.DB
	err := db.AutoMigrate(&model.Books{})
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepo(db)
	app := service.NewService(repo)
	route.RegisterApi(router, app)

	port := os.Getenv("PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
