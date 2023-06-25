package userservice

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang_web_Project/database"
	"golang_web_Project/model"
	"log"
	// "time"
)

func GetUser() []model.User {

	users := []model.User{}

	db := database.MySqlConnection()
	rows, err := db.Query("select * from users")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	// Iterate over the query results
	for rows.Next() {
		var user model.User
		var _createdAt string
		var _updatedAt sql.NullString

		err := rows.Scan(&user.Id, &user.Nama, &user.Email, &user.Password, &_createdAt, &_updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		// if _updatedAt.Valid {
		// 	user.Updated_at = _updatedAt
		// }

		// value, err := time.Parse("", _createdAt)
		// user.Created_at = value

		users = append(users, user)

	}

	// fmt.Println(users)

	return users
}

func AddUser(data *model.User) {
	db := database.MySqlConnection()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	stmt, err := tx.Prepare("INSERT INTO users (nama, email, password) VALUES (?, ?, ?)")
	if err != nil {
		tx.Rollback()
	}

	defer stmt.Close()

	// Execute the SQL statement with the user data
	_, err = stmt.Exec(data.Nama, data.Email, data.Password)
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}
	defer stmt.Close()
	defer db.Close()

}
