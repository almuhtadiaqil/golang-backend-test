package handlers

import (
	"backend-test/src/entities"
	request "backend-test/src/request/product"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductService interface {
	Create(ctx *gin.Context, payload request.StoreProductRequest) (response entities.Product, err error)
	Retrieved(paginateRequest request.ProductPaginationRequest) (*entities.PaginateResponse, error)
}

func GetPagincation(service ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var queryParams request.ProductPaginationRequest
		var response entities.BaseResponse

		if err := c.ShouldBind(&queryParams); err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()
			c.JSON(response.Code, response)
			return
		}

		responses, err := service.Retrieved(queryParams)
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			c.JSON(response.Code, response)
			return
		}

		c.JSON(responses.Code, responses)
		return
	}
}

func StoreProduct(service ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload request.StoreProductRequest
		var response entities.BaseResponse

		validate := validator.New()

		if err := c.BindJSON(&payload); err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()
			c.JSON(response.Code, response)
			return
		}

		if err := validate.Struct(&payload); err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()
			c.JSON(response.Code, response)
			return
		}

		product, err := service.Create(c, payload)
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			c.JSON(response.Code, response)
			return
		}
		response.Code = http.StatusCreated
		response.Data = product
		response.Message = "Successfully created a new product."

		c.JSON(response.Code, response)
		return
	}
}
