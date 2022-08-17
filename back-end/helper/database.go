package helper

import (
	"database/sql"
)

func Migrate(DB *sql.DB) {
	//CODE
	pass1 := "apahayo"
	pass2 := "123456"
	sqlStatement := `
		INSERT INTO user (name, email, role, password)
		VALUES 
			('Arief', 'coco@gmail.com', "penjual", ?),
			('Sapacikk', 'apa@gmail.com', "pembeli", ?);
		
		INSERT INTO product (name, price, desc, cur_quantity, quantity, id_owner)
		VALUES
			('Harga Diri Guwe', 3000, 'waodaowjdoiawjdwaijdoawjdoiajwdowaijdoaiwjdowajdo', 3, 3, 1),
			('Batu Dari Gunung Gunungan',1000, 'iwjadiwjadojawdoajwdoiwajdowaijdoawjwdaoij', 100, 100, 1),
			('Tisu Putih Bekas Cebok', 5000, 'wjadoiajwdoaiwjdoiawdoaiwjdwauwahdoaoijdawjdiowajdoiaw', 20, 30, 1);
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
		PRAGMA foreign_keys = ON;
		CREATE TABLE IF NOT EXISTS user (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL,
			role TEXT NOT NULL,
			image_loc TEXT,
			UNIQUE(email)
		);
		CREATE TABLE IF NOT EXISTS product (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			price INTEGER NOT NULL,
			desc TEXT NOT NULL,
			cur_quantity INTEGER NOT NULL,
			quantity INTEGER NOT NULL,
			image_loc TEXT,
			id_owner INTEGER,
			FOREIGN KEY (id_owner)
			REFERENCES user (id)
				ON UPDATE CASCADE
				ON DELETE SET NULL
		);
		CREATE TABLE IF NOT EXISTS order_record (
			amount INTEGER NOT NULL,
			id_buyer INTEGER,
			id_product INTEGER,
			FOREIGN KEY (id_buyer)
			REFERENCES user (id)
				ON UPDATE CASCADE
				ON DELETE SET NULL,
			FOREIGN KEY (id_product)
			REFERENCES product (id)
				ON UPDATE CASCADE
				ON DELETE SET NULL
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
