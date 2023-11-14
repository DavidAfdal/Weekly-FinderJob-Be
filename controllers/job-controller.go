package controller

import (
	"fmt"

	helper "github.com/DavidAfdal/Weekly-FinderJob-Be/helpers"
	"github.com/DavidAfdal/Weekly-FinderJob-Be/models/request"
	"github.com/DavidAfdal/Weekly-FinderJob-Be/models/response"
	usecase "github.com/DavidAfdal/Weekly-FinderJob-Be/usecases"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type JobController struct {
	jobUseCase *usecase.JobUseCase
	Validate       *validator.Validate
}

func NewJobController(jobUseCase *usecase.JobUseCase, Validate *validator.Validate) *JobController {
	return &JobController{
		jobUseCase: jobUseCase,
		Validate: Validate,
	}
}

func (c *JobController) GetJobs(ctx *fiber.Ctx) error {
	jobResponse := c.jobUseCase.FindAll()
  
	response := response.WebResponse{
	   Message: "Sucess Get All Data",
	   Status: "OK",
	   Data: jobResponse,
	}
  
	return ctx.Status(200).JSON(response)
  }


  func (c *JobController) CreateJob(ctx *fiber.Ctx) error{
	var JobRequest request.CreateJobInput

	err := ctx.BodyParser(&JobRequest)

	if err != nil {
		helper.ErrorPanic(err)
	}

	errValidate := c.Validate.Struct(JobRequest)

	if errBindResult := helper.BindErrorHandler(errValidate); errBindResult != nil {
	   errorResponse := response.ErrorResponse{
		   Status: "Bad Request",
		   Errors: errBindResult,
	   }
	   return ctx.Status(400).JSON(errorResponse)

   }

	c.jobUseCase.Create(JobRequest)

	webResponse := response.WebResponse{
	   Message: "Succes add data job",
	   Status: "Created",
	   Data: nil,
	}

	return ctx.Status(201).JSON(webResponse)
}


func (c *JobController) GetJobByCategory(ctx *fiber.Ctx) error{
	jobCategory := ctx.Query("category")
	
	fmt.Println(jobCategory)

	jobResponse := c.jobUseCase.FindByCategory(jobCategory)


	response := response.WebResponse{
		Message: "Success Get Data Job By " + jobCategory,
		Status: "Ok",
		Data: jobResponse,
	}

	return ctx.Status(200).JSON(response)
}

func (c *JobController) GetJobByUserId(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId") 
	

	jobResponse := c.jobUseCase.FindByUserId(userId)

	response := response.WebResponse{
		Message: "Success Get Data Job By UserId",
		Status: "Ok",
		Data: jobResponse,
	}
    return ctx.Status(200).JSON(response)
}


func (c *JobController) GetJobById(ctx *fiber.Ctx) error{
	jobId, err:=  ctx.ParamsInt("jobId")
	helper.ErrorPanic(err)

	jobResponse, err := c.jobUseCase.FindById(jobId)

	if err != nil {
		fmt.Println(err)
		errRepsonse := response.ErrorResponse{
			Errors: err.Error(),
			Status: "Not Found",
		}
		return ctx.Status(404).JSON(errRepsonse)
		 
	}
	response := response.WebResponse{
		Message: "Succes Get Data By Id",
		Status: "Ok",
		Data: jobResponse,
	}

	return ctx.Status(200).JSON(response)
}