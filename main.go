package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Konfigurasi koneksi MySQL
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/mahasiswa")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Membaca data dari tabel
	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Menampilkan data
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
}
