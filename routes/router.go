package routes

import (
	controller "github.com/DavidAfdal/Weekly-FinderJob-Be/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, jobController *controller.JobController, applierController *controller.ApplierController) {
    api := app.Group("/api")
   
        // Setup the Node Routes
    SetupJobRoutes(api, jobController)
    SetupApplierRoutes(api, applierController)
}