package userservice

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang_web_Project/database"
	"log"
	"time"
)

type User struct {
	Id         int
	Nama       string
	Email      string
	Password   string
	Created_at time.Time
	Updated_at sql.NullString
}

func GetUser() []User {

	users := []User{}

	db := database.MySqlConnection()
	rows, err := db.Query("select * from users")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	// Iterate over the query results
	for rows.Next() {
		var user User
		var _createdAt string
		var _updatedAt sql.NullString

		err := rows.Scan(&user.Id, &user.Nama, &user.Email, &user.Password, &_createdAt, &_updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		if _updatedAt.Valid {
			user.Updated_at = _updatedAt
		}

		value, err := time.Parse("2006-01-02 15:04:05", _createdAt)
		user.Created_at = value

		users = append(users, user)

	}

	// fmt.Println(users)

	return users
}
