package models

type Product struct {
	ID       uint   `db:"id"`
	Name     string `db:"name"`
	Price    uint   `db:"price"`
	ImageLoc string `db:"image_loc"`
}

type Users struct {
	ID       uint   `db:"id"`
	Name     string `db:"name"`
	Password string `db:"price"`
	Email    string `db:"email"`
	ImageLoc string `db:"image_loc"`
}
