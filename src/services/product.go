package services

import (
	"backend-test/src/entities"
	request "backend-test/src/request/product"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type ProductService struct {
	repository  ProductRepository
	redisClient *redis.Client
}

type ProductRepository interface {
	Retrieved(props request.ProductPaginationRequest) (products []entities.Product, total int64, err error)
	Create(product *entities.Product) (err error)
}

func NewProductService(repository ProductRepository, redisClient *redis.Client) *ProductService {
	return &ProductService{
		repository:  repository,
		redisClient: redisClient,
	}
}

func (s *ProductService) Create(ctx *gin.Context, payload request.StoreProductRequest) (product entities.Product, err error) {
	product.ProductCode = payload.ProductCode
	product.ProductName = payload.ProductName
	product.ProductPrice = payload.ProductPrice
	product.ProductQty = payload.ProductQty
	product.ProductPrice = payload.ProductPrice

	s.redisClient.Set(ctx, fmt.Sprintf("product_%s", product.ProductCode), product, 24*time.Hour)
	// config.SaveDataFromRedisToDB(ctx)
	err = s.repository.Create(&product)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *ProductService) Retrieved(paginateRequest request.ProductPaginationRequest) (*entities.PaginateResponse, error) {
	var response entities.PaginateResponse
	var productsResponses []entities.ProductResponse

	if paginateRequest.PageIndex <= 0 {
		paginateRequest.PageIndex = 1
	}

	if paginateRequest.PageSize <= 0 {
		paginateRequest.PageSize = 25
	}

	products, total, err := s.repository.Retrieved(paginateRequest)
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		productResponse := product.FormatAsResponse()
		productsResponses = append(productsResponses, *productResponse)
	}

	response.Total = total
	response.Data = productsResponses
	response.PageIndex = paginateRequest.PageIndex
	response.PageSize = paginateRequest.PageSize
	response.Message = "Successfull to retrieved data"
	response.Code = http.StatusOK

	return &response, nil
}
