package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func MySqlConnection() *sql.DB {
	username := "root"
	password := "root"
	host := "localhost"
	port := "8889"
	database := "layar_tancep"

	// Create the MySQL connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	// open a connection to the database
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MySQL database!")

	if err != nil {
		log.Fatal(err)
	}
	return db
}
