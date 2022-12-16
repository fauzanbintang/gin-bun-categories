package routers

import (
	"net/http"
	"zamannow/go-rest-api/controllers"
	"zamannow/go-rest-api/domain/repository"
	"zamannow/go-rest-api/dto/responses"
	"zamannow/go-rest-api/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RouteCallback func(r *gin.RouterGroup) *gin.RouterGroup

type GinEng gin.Engine

func RoutesHandler(r *gin.Engine) {
	rp := repository.InitRepositoryInstance()
	validate := validator.New()

	categorySrv := services.NewCategoryService(rp.Category)
	categoryCtl := controllers.NewCategoryController(categorySrv, validate)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, responses.R{
			Code:    http.StatusOK,
			Message: "Hello World :)",
			Data: gin.H{
				"connected": true,
			},
		})
	})

	category := r.Group("/categories")
	{
		category.GET("", categoryCtl.FindAll)
		category.POST("", categoryCtl.Create)
		category.GET(":id", categoryCtl.FindById)
		category.PUT(":id", categoryCtl.Update)
		category.DELETE(":id", categoryCtl.Delete)

	}
}
