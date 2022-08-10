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
	sqlStatement := `
	INSERT INTO user (name, password, email, role, image_loc) 
		VALUES (?, ?, ?, ?, ?);`

	_, err := r.db.Exec(sqlStatement, user.Name, user.Password, user.Email, user.Role, user.ImageLoc)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *UserRepository) DeleteUser(id int) error {
	sqlstmt := `DELETE FROM user WHERE id = ?;`

	_, err := r.db.Exec(sqlstmt, id)

	if err != nil {
		log.Println("Error When Deleting in a Row : ", err.Error())
		return err
	}
	return nil
}

func (r *UserRepository) UpdateUser(user models.Users) error {
	sqlStatement := `
		UPDATE user
		SET name = ?,
			email = ?,
			image_loc = ?
		WHERE 
			id = ?;
	`

	_, err := r.db.Exec(sqlStatement, user.Name, user.Email, user.ImageLoc, user.ID)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *UserRepository) UpdateUserPassword(id int, hashedPassword string) error {

	sqlStatement := `
		UPDATE user
		SET password = ?
		WHERE 
			id = ?;
	`
	_, err := r.db.Exec(sqlStatement, hashedPassword, id)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *UserRepository) GetUsers() ([]models.Users, error) {
	sqlStatement := `SELECT * FROM user;`

	rows, err := r.db.Query(sqlStatement)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := []models.Users{}

	for rows.Next() {
		var u models.Users
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role, &u.ImageLoc)
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

	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role, &u.ImageLoc)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.Users, error) {
	sqlStatement := "SELECT * FROM user WHERE email = ?"
	row := r.db.QueryRow(sqlStatement, email)

	var u models.Users

	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role, &u.ImageLoc)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &u, nil
}
