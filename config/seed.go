package config

import (
	"backend-test/src/entities"
	"math/rand"

	"github.com/go-faker/faker/v4"

	"gorm.io/gorm"
)

func DBSeed(db *gorm.DB) {
	productTypes := []string{"sayuran", "protein", "buah", "snack"}
	productPrices := []float32{10000.00, 15000.00, 5000.00, 1000.00}

	var count int64
	db.Model(entities.Product{}).Count(&count)
	if count <= 1 {
		for i := 0; i <= 30; i++ {
			product := entities.Product{
				ProductName:  faker.Word(),
				ProductCode:  faker.Word(),
				ProductType:  productTypes[rand.Intn(len(productTypes))],
				ProductPrice: productPrices[rand.Intn(len(productPrices))],
				ProductQty:   rand.Intn(100),
			}
			db.Save(&product)
		}
	}
}
