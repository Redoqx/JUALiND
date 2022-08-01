package models

import "database/sql"

type Product struct {
	ID       uint           `db:"id"`
	Name     string         `db:"name"`
	Price    uint           `db:"price"`
	ImageLoc sql.NullString `db:"image_loc"`
}

type Users struct {
	ID       uint           `db:"id"`
	Name     string         `db:"name"`
	Password string         `db:"password"`
	Email    string         `db:"email"`
	ImageLoc sql.NullString `db:"image_loc"`
}
