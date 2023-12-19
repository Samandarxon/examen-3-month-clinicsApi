package main

import (
	"log"

	"github.com/Samandarxon/examen_3-month/clinics/api"
	"github.com/Samandarxon/examen_3-month/clinics/config"
	_ "github.com/Samandarxon/examen_3-month/clinics/migrations"
	"github.com/Samandarxon/examen_3-month/clinics/storage/postgres"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()
	pgStorage, err := postgres.NewConnectionPostgres(&cfg)
	if err != nil {
		panic(err)
	}
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	api.SetUpAPI(r, &cfg, pgStorage)

	log.Println("Listening:", cfg.ServiceHost+cfg.ServiceHTTPPort, "...")
	log.Println("Swagger: http://"+cfg.ServiceHost+cfg.ServiceHTTPPort+"/swagger/index.html", "...")

	if err := r.Run(cfg.ServiceHost + cfg.ServiceHTTPPort); err != nil {
		panic("Listent and service panic:" + err.Error())
	}
}
