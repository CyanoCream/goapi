package app

import (
	"fmt"
	"os"
	"sesi_8/config"
	"sesi_8/repository"
	"sesi_8/route"
	"sesi_8/service"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.PSQL.DB)
	app := service.NewService(repo)
	route.RegisterApi(router, app)

	port := os.Getenv("PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
