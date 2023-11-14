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

type ApplierController struct {
	ApplierUseCase *usecase.ApplierUseCase
	Validate       *validator.Validate
}

func NewApplierController(ApplierUseCase *usecase.ApplierUseCase, Validate       *validator.Validate) *ApplierController {
	return &ApplierController{
		ApplierUseCase: ApplierUseCase,
		Validate: Validate,
	}
}

func (c *ApplierController) ApplyJob(ctx *fiber.Ctx) error{
	var ApplierRequest request.ApplierRequest

	err := ctx.BodyParser(&ApplierRequest)

	if err != nil {
		helper.ErrorPanic(err)
	}

	errValidate := c.Validate.Struct(ApplierRequest)
	
	if errBindResult := helper.BindErrorHandler(errValidate); errBindResult != nil {
		errorResponse := response.ErrorResponse{
			Status: "Bad Request",
			Errors: errBindResult,
		}
		return ctx.Status(400).JSON(errorResponse)
 
	}

	errorApply := c.ApplierUseCase.ApplyJob(ApplierRequest)

	fmt.Println(errorApply)

	if errorApply != nil {
		errRepsonse := response.ErrorResponse{
			Errors: errorApply.Error(),
			Status: "Bad Request",
		}
		return ctx.Status(400).JSON(errRepsonse)
	}

	webResponse := response.WebResponse{
		Message: "Berhasil Apply Job",
		Status:  "Ok",
		Data:    nil,
	}

   return ctx.Status(200).JSON(webResponse)
}

func (c *ApplierController) GetByUserId(ctx *fiber.Ctx) error {
	UserId := ctx.Query("userId")

	fmt.Println(UserId)

	jobResponse, err := c.ApplierUseCase.FindByUserId(UserId)

	if err != nil {
		fmt.Println(err)
		errRepsonse := response.ErrorResponse{
			Errors: err.Error(),
			Status: "Not Found",
		}
		return ctx.Status(400).JSON(errRepsonse)
	
	}

	response := response.WebResponse{
		Message: "Success ambil data job berdasarkan user Id ",
		Status:  "Ok",
		Data:    jobResponse,
	}

	 return ctx.Status(200).JSON(response)
}