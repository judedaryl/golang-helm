package product

import (
	"time"

	"optumlabs.com/myapi/internal/entities"
)

type Service interface {
	Create(req CreateProductRequest) (Product, error)
	Get(id string) (Product, error)
	GetAll() ([]Product, error)
}

type Product struct {
	entities.Product
}

type CreateProductRequest struct {
	Name string `json:"name"`
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return service{repo}
}

func (s service) Create(req CreateProductRequest) (Product, error) {
	id := entities.GenerateID()
	now := time.Now()
	product, err := s.repo.Create(entities.Product{
		ID:        id,
		Name:      req.Name,
		CreatedAt: now,
	})

	if err == nil {
		return Product{product}, err
	}
	return Product{}, err
}

func (s service) Get(id string) (Product, error) {
	product, err := s.repo.Get(id)
	if err == nil {
		return Product{product}, err
	}
	return Product{}, err
}

func (s service) GetAll() ([]Product, error) {
	result := []Product{}
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	for _, item := range products {
		result = append(result, Product{item})
	}
	return result, err
}
