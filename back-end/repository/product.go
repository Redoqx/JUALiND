package repository

import (
	"JUALiND/models"
	"database/sql"
	"fmt"
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
	sqlStatement := `
	INSERT INTO product (name, price, desc, cur_quantity, quantity, image_loc, id_owner) 
	VALUES (?, ?, ?, ?, ?, ?, ?);`

	_, err := r.db.Exec(sqlStatement, product.Name, product.Price, product.Description, product.CurrentQuantity, product.Quantity, product.ImageLoc, product.OwnerID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *ProductRepository) UpdateProduct(product models.Product) error {
	sqlStatement := `
		UPDATE product
		SET name = ?,
			price = ?,
			desc =?,
			cur_quantity = ?,
			quantity = ?,
			image_loc = ?
		WHERE 
			id = ?;
	`

	_, err := r.db.Exec(sqlStatement, product.Name, product.Price, product.Description, product.CurrentQuantity, product.Quantity, product.ImageLoc, product.ID)

	if err != nil {
		log.Println(err)
		return err
	}

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

	err := row.Scan(&u.ID, &u.Name, &u.Price, &u.Description, &u.CurrentQuantity, &u.Quantity, &u.ImageLoc, &u.OwnerID)

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
		err := rows.Scan(&u.ID, &u.Name, &u.Price, &u.Description, &u.CurrentQuantity, &u.Quantity, &u.ImageLoc, &u.OwnerID)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, u)
	}
	return result, nil
}

func (r *ProductRepository) GetProductByUser(userID int) ([]models.Product, error) {
	sqlStatement := `
		SELECT * FROM product WHERE id_owner = ?;
	`

	rows, err := r.db.Query(sqlStatement, userID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var p []models.Product

	for rows.Next() {
		var item models.Product
		err = rows.Scan(&item.ID, &item.Name, &item.Price, &item.Description, &item.CurrentQuantity, &item.Quantity, &item.ImageLoc, &item.OwnerID)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		p = append(p, item)
	}

	return p, nil

}

func (r *ProductRepository) GetProductByName(name string) ([]models.Product, error) {
	sqlStatement := `
		SELECT * FROM product WHERE name LIKE ?;
	`
	pattern := fmt.Sprintf("%c%s%c", '%', name, '%')
	rows, err := r.db.Query(sqlStatement, pattern)
	log.Println("Pattern :", pattern)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var p []models.Product

	for rows.Next() {
		var item models.Product
		err = rows.Scan(&item.ID, &item.Name, &item.Price, &item.Description, &item.CurrentQuantity, &item.Quantity, &item.ImageLoc, &item.OwnerID)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		p = append(p, item)
	}

	return p, nil

}
