package repository

import (
	"JUALiND/models"
	"JUALiND/schema"
	"database/sql"
	"log"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}
func (o *OrderRepository) CreateOrder(order models.Order) error {
	sqlStatement := `INSERT INTO order_record (amount, id_buyer, id_product, date) VALUES (?, ?, ?, date('now'))`

	_, err := o.db.Exec(sqlStatement, order.Amount, order.BuyerID, order.ProductID)
	if err != nil {
		return err
	}

	return nil

}
func (o *OrderRepository) GetAllOrdersByUser(ownerID int) ([]schema.OrderStruct, error) {
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

	var results []schema.OrderStruct

	for rows.Next() {
		var o schema.OrderStruct
		err = rows.Scan(&o.BuyerName, &o.ProductName, &o.ProductPrice, &o.Amount, &o.Date, &o.ConfirmationLink)

		if err != nil {
			log.Println("Error : ", err.Error())
			return nil, err
		}
		results = append(results, o)
	}

	return results, nil
}
