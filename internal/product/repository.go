package product

import (
	"optumlabs.com/myapi/internal/entities"
	"optumlabs.com/myapi/pkg/dbcontext"
)

type Repository interface {
	Create(product entities.Product) (entities.Product, error)
	Get(id string) (entities.Product, error)
	GetAll() ([]entities.Product, error)
}
type repository struct {
	db *dbcontext.DB
}

func NewRepository(db *dbcontext.DB) Repository {
	return &repository{db}
}

func (r repository) Create(product entities.Product) (entities.Product, error) {
	result := r.db.DB().Create(&product)
	return product, result.Error
}

func (r repository) Get(id string) (entities.Product, error) {
	var product entities.Product
	result := r.db.DB().First(&product, "id = ?", id)
	return product, result.Error
}

func (r repository) GetAll() ([]entities.Product, error) {
	var products []entities.Product
	result := r.db.DB().Find(&products)
	return products, result.Error
}
