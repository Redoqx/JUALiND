package repository

import (
	"JUALiND/models"
	"database/sql"
	"log"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
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
	sqlstmt := `DELETE FROM product WHERE id = ?`

	_, err := r.db.Exec(sqlstmt, id)

	if err != nil {
		log.Println("Error When Deleting in a Row : ", err.Error())
		return err
	}
	return nil
}

func (r *ProductRepository) GetProduct(id int) (*models.Product, error) {
	sqlStatement := `SELECT * FROM product WHERE id = ?`

	row := r.db.QueryRow(sqlStatement, id)

	var u models.Product

	err := row.Scan(&u.ID, &u.Name, &u.Price, &u.ImageLoc)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &u, nil
}

func (r *ProductRepository) GetProducts() ([]models.Product, error) {
	sqlStatement := `SELECT * FROM product`

	rows, err := r.db.Query(sqlStatement)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := []models.Product{}

	for rows.Next() {
		var u models.Product
		err := rows.Scan(&u.ID, &u.Name, &u.Price, &u.ImageLoc)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, u)
	}
	return result, nil
}