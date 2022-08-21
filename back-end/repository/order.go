package repository

import (
	"database/sql"
	"log"
)

type OrderRepository struct {
	db *sql.DB
}

type OrderStruct struct {
	BuyerName        string         `json:"buyer_name"`
	ProductName      string         `json:"product_name"`
	ProductPrice     int            `json:"product_price"`
	Amount           int            `json:"amount"`
	Date             string         `json:"date"`
	ConfirmationLink sql.NullString `json:"confirmation_link"`
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (o *OrderRepository) GetAllOrdersByUser(ownerID int) ([]OrderStruct, error) {
	sqlStatement := `
		SELECT b.name, p.name, p.price, o.amount, o.date, o.confirmation_link 
		FROM 
			order_record as o
		INNER JOIN product as p ON o.id_product = p.id
		INNER JOIN user as b ON o.id_buyer = b.id
		WHERE p.id_owner = ?; 
	`

	rows, err := o.db.Query(sqlStatement, ownerID)

	if err != nil {
		log.Println("Error : ", err.Error())
		return nil, err
	}

	var results []OrderStruct

	for rows.Next() {
		var o OrderStruct
		err = rows.Scan(&o.BuyerName, &o.ProductName, &o.ProductPrice, &o.Amount, &o.Date, &o.ConfirmationLink)

		if err != nil {
			log.Println("Error : ", err.Error())
			return nil, err
		}
		results = append(results, o)
	}

	return results, nil
}
