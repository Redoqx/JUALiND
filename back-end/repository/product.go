package repository

import (
	"JUALiND/models"
	"database/sql"
)

type ProductRepository struct {
	db *sql.DB
}

func (r *ProductRepository) NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) CreateProduct(product models.Product) error {
	return nil
}

func (r *ProductRepository) UpdateProduct(product models.Product) error {
	return nil
}

func (r *ProductRepository) DeleteProduct(id int) error {
	return nil
}

func (r *ProductRepository) GetProduct() (*models.Product, error) {
	return nil, nil
}

func (r *ProductRepository) GetProducts() ([]models.Product, error) {
	return nil, nil
}
