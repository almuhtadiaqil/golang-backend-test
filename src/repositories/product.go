package repositories

import (
	"backend-test/src/entities"
	request "backend-test/src/request/product"
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (m ProductRepository) Retrieved(props request.ProductPaginationRequest) (products []entities.Product, total int64, err error) {
	query := m.db.Debug().Model(&products)

	if props.Query != "" {
		if id, err := strconv.Atoi(props.Query); err == nil {
			query = query.Where("product_id = ?", id)
		}
		search := "%" + strings.ToLower(props.Query) + "%"
		query = query.Or("lower(product_name) like ?", search)
	}

	if _, ok := props.Filter["type"]; ok {
		productTypes := strings.Split(props.Filter["type"].(string), ",")
		query = query.Where("product_type in ?", productTypes)
	}

	if props.SortBy != "" {
		orderBy := strings.Split(props.SortBy, ",")
		if len(orderBy) > 1 {
			query = query.Order(fmt.Sprintf("%s %s", orderBy[0], orderBy[1]))
		} else {
			query = query.Order(fmt.Sprintf("%s ASC", orderBy[0]))
		}
	}

	query.Count(&total)

	offset := (props.PageIndex - 1) * props.PageSize

	query = query.Offset(offset).Limit(props.PageSize)

	err = query.Find(&products).Error

	return products, total, err
}

func (m ProductRepository) Create(product *entities.Product) (err error) {
	if err = m.db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (m ProductRepository) Detail(product *entities.Product, id int) (err error) {
	if err = m.db.First(&product, "product_id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (m ProductRepository) Update(data *entities.Product, id int) (err error) {
	var product entities.Product

	if err = m.Detail(&product, id); err != nil {
		return err
	}

	err = m.db.Updates(data).Where("product_id = ?", id).Error

	return err
}

func (m ProductRepository) Delete(id int) (err error) {
	var product entities.Product

	if err = m.Detail(&product, id); err != nil {
		return err
	}

	if err = m.db.Delete(&product).Error; err != nil {
		return err
	}

	return nil
}
