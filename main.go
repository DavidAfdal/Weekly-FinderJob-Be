package main

import (
	"github.com/DavidAfdal/Weekly-FinderJob-Be/config"
	controller "github.com/DavidAfdal/Weekly-FinderJob-Be/controllers"
	repository "github.com/DavidAfdal/Weekly-FinderJob-Be/repositorys"
	"github.com/DavidAfdal/Weekly-FinderJob-Be/routes"
	usecase "github.com/DavidAfdal/Weekly-FinderJob-Be/usecases"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
    viper.ReadInConfig()
}

func main() {

	app := fiber.New()
	app.Use(cors.New())

	DB := config.ConnectDB()
	validate := validator.New()

	jobRepo := repository.NewJobRepository(DB)
	jobUseCase := usecase.NewJobsUseCase(jobRepo)
	jobController := controller.NewJobController(jobUseCase, validate)
	applierRepo := repository.NewApplierRepository(DB)
	applierUseCase := usecase.NewApplierUseCase(applierRepo, jobRepo)
	applierController := controller.NewApplierController(applierUseCase, validate)
    

	routes.SetupRoutes(app, jobController, applierController)





    // Send a string back for GET calls to the endpoint "/"


    // Listen on PORT 3000
    app.Listen(":5000")
}