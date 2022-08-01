package helper

import (
	"database/sql"
)

func Migrate(DB *sql.DB) {
	//CODE
	pass1 := "apahayo"
	pass2 := "123456"
	sqlStatement := `
		INSERT INTO user (name, email, password)
		VALUES 
			('Arief', 'coco@gmail.com', ?),
			('Sapacikk', 'apa@gmail.com', ?);
		
		INSERT INTO product (name, price)
		VALUES
			('Harga Diri Guwe', 2000),
			('Batu Dari Gunung Gunungan', 1000),
			('Tisu Putih Bekas Cebok', 5000);
		`

	hash1, err := HashPassword(pass1)
	hash2, err := HashPassword(pass2)

	if err != nil {
		panic(err)
	}

	_, err = DB.Exec(sqlStatement, hash2, hash1)
	if err != nil {
		panic(err)
	}
}

func InitDB(DB *sql.DB) {
	sqlStatement := `
		CREATE TABLE IF NOT EXISTS product (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			price TEXT NOT NULL,
			image_loc TEXT
		);

		CREATE TABLE IF NOT EXISTS user (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL,
			image_loc TEXT,
			UNIQUE(email)
		);
	`

	_, err := DB.Exec(sqlStatement)

	if err != nil {
		panic(err)
	}

}

func NullStringToString(s sql.NullString) string {
	if s.Valid {
		return s.String
	} else {
		return ""
	}
}

func StringToNullString(s string) sql.NullString {
	if len(s) < 1 {
		return sql.NullString{
			Valid: false,
		}
	} else {
		return sql.NullString{
			String: s,
			Valid:  true,
		}
	}
}
