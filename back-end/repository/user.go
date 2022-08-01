package repository

import (
	"JUALiND/models"
	"database/sql"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
func (r *UserRepository) CreateUser(user models.Users) error {
	return nil
}

func (r *UserRepository) DeleteUser(id int) error {
	sqlstmt := `DELETE FROM user WHERE id = ?`

	_, err := r.db.Exec(sqlstmt, id)

	if err != nil {
		log.Println("Error When Deleting in a Row : ", err.Error())
		return err
	}
	return nil
}

func (r *UserRepository) UpdateUser(user models.Users) error {
	return nil
}

func (r *UserRepository) GetUsers() ([]models.Users, error) {
	sqlStatement := `SELECT * FROM user`

	rows, err := r.db.Query(sqlStatement)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := []models.Users{}

	for rows.Next() {
		var u models.Users
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.ImageLoc)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, u)
	}
	return result, nil

}

func (r *UserRepository) GetUser(id int) (*models.Users, error) {
	sqlStatement := `SELECT * FROM user WHERE id = ?`

	row := r.db.QueryRow(sqlStatement, id)

	var u models.Users

	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.ImageLoc)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &u, nil
}
