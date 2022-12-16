package controllers

import (
	"net/http"
	"strconv"
	"zamannow/go-rest-api/dto/requests"
	"zamannow/go-rest-api/dto/responses"
	"zamannow/go-rest-api/services"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type CategoryController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type categoryController struct {
	categorySrv services.CategoryService
	validate    *validator.Validate
}

func NewCategoryController(categorySrv services.CategoryService, validate *validator.Validate) CategoryController {
	return &categoryController{
		categorySrv: categorySrv,
		validate:    validate,
	}
}

func (ctl *categoryController) Create(ctx *gin.Context) {
	var src requests.CreateCategoryRequest

	if err := ctl.ParseRequestEntity(ctx, &src); err != nil {
		panic(err)
	}

	res, err := ctl.categorySrv.Create(ctx, src)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusCreated, responses.R{
		Code:    http.StatusCreated,
		Message: "Succcessfully create Category",
		Data:    res,
	})
}

func (ctl *categoryController) Update(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 1<<6)
	if err != nil {
		panic(err)
	}

	var src requests.UpdateCategoryRequest
	src.ID = id
	if err := ctl.ParseRequestEntityUpdate(ctx, &src); err != nil {
		panic(err)
	}

	res, err := ctl.categorySrv.Update(ctx, src)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, responses.R{
		Code:    http.StatusOK,
		Message: "Successfully update data",
		Data:    res,
	})
}

func (ctl *categoryController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 1<<6)
	if err != nil {
		panic(err)
	}

	err = ctl.categorySrv.Delete(ctx, id)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, responses.R{
		Code:    http.StatusOK,
		Message: "Successfully delete data",
	})
}

func (ctl *categoryController) FindById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 1<<6)
	if err != nil {
		panic(err)
	}

	res, err := ctl.categorySrv.FindById(ctx, id)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, responses.R{
		Code:    http.StatusOK,
		Message: "Successfully get data",
		Data:    res,
	})
}

func (ctl *categoryController) FindAll(ctx *gin.Context) {
	res, err := ctl.categorySrv.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, responses.R{
		Code:    http.StatusOK,
		Message: "Successfully get all data",
		Data:    res,
	})
}

func (ctl *categoryController) ParseRequestEntity(ctx *gin.Context, src *requests.CreateCategoryRequest) error {
	if err := ctx.ShouldBindBodyWith(src, binding.JSON); err != nil {
		panic(err)
	}
	if err := ctl.validate.Struct(src); err != nil {
		panic(err)
	}
	return nil
}

func (ctl *categoryController) ParseRequestEntityUpdate(ctx *gin.Context, src *requests.UpdateCategoryRequest) error {
	if err := ctx.ShouldBindBodyWith(src, binding.JSON); err != nil {
		panic(err)
	}
	if err := ctl.validate.Struct(src); err != nil {
		panic(err)
	}
	return nil
}
