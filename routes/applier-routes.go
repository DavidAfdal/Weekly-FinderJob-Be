package routes

import (
	controller "github.com/DavidAfdal/Weekly-FinderJob-Be/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupApplierRoutes(router fiber.Router, applierController *controller.ApplierController) {
	applier := router.Group("/applier")
	

	applier.Post("/apply", applierController.ApplyJob)
}