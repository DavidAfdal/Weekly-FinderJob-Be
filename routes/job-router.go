package routes

import (
	controller "github.com/DavidAfdal/Weekly-FinderJob-Be/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupJobRoutes(router fiber.Router, jobController *controller.JobController) {
	job := router.Group("/job")
	// Create a Note
	job.Get("/", jobController.GetJobs)
	job.Get("/filter", jobController.GetJobByCategory)
	job.Get("/:jobId", jobController.GetJobById)
	job.Get("/created/:userId", jobController.GetJobByUserId)
	job.Post("/", jobController.CreateJob)
	// Read all Notes
	// // Read one Note
	// note.Get("/:noteId", noteHandler.GetNote)
	// // // Update one Note
	// note.Put("/:noteId", noteHandler.UpdateNote)
	// // // Delete one Note
	// note.Delete("/:noteId", noteHandler.DeleteNote)
}