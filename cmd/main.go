package main

import (
	"log"
	"os"

	"gitlab.com/p9359/backend-prob/febry-go/api"
	"gitlab.com/p9359/backend-prob/febry-go/internal/app"
	"gitlab.com/p9359/backend-prob/febry-go/internal/repository"
	"gitlab.com/p9359/backend-prob/febry-go/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Unable to load .env file: %v", err)
	}
}

func main() {
	mode := os.Getenv("GIN_MODE")
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	sqlx, err := repository.NewSQLDB()
	if err != nil {
		log.Printf("err: %v", err)
		return
	}

	repository := repository.NewDAO(nil, sqlx)
	service := service.NewReviewService(repository)
	app := app.NewReviewApp(service)

	router := gin.Default()

	// register routes
	api.RegisterRoutes(router, *app)

	router.Run(":9080")
}
