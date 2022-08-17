package models

import "database/sql"

type Product struct {
	ID              uint           `db:"id"`
	OwnerID         uint           `db:"owner_id"`
	Name            string         `db:"name"`
	Price           uint           `db:"price"`
	Description     string         `db:"desc"`
	CurrentQuantity uint           `db:"cur_quantity"`
	Quantity        uint           `db:"quantity"`
	ImageLoc        sql.NullString `db:"image_loc"`
}

type Users struct {
	ID       uint           `db:"id"`
	Name     string         `db:"name"`
	Password string         `db:"password"`
	Email    string         `db:"email"`
	Role     string         `db:"role"`
	ImageLoc sql.NullString `db:"image_loc"`
}

type Order struct {
	Amount    uint `db:"amount"`
	ProductID uint `db:"product_id"`
	BuyerID   uint `db:"user_id"`
}
