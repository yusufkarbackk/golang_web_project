package userservice

import (
	"context"
	"golang_web_Project/database"
	"golang_web_Project/model"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/api/option"
	// "time"
)

func GetUsers() []model.UserNoPassword {

	users := []model.UserNoPassword{}

	db := database.MySqlConnection()
	rows, err := db.Query("select uuid, nik, nama, jenis_kelamin, alamat, role, saldo from users")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	// Iterate over the query results
	for rows.Next() {
		var user model.UserNoPassword

		err := rows.Scan(&user.Uuid, &user.Nik, &user.Nama, &user.Jenis_kelamin, &user.Alamat, &user.Role, &user.Saldo)
		if err != nil {
			log.Fatal(err)
		}

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

	stmt, err := tx.Prepare("INSERT INTO users (nik, nama, jenis_kelamin, alamat, password, saldo) VALUES (?, ?, ?, ?, ?, 0)")
	if err != nil {
		tx.Rollback()
	}

	defer stmt.Close()

	// Execute the SQL statement with the user data
	_, err = stmt.Exec(data.Nik, data.Nama, data.Jenis_kelamin, data.Alamat, data.Password)
	if err != nil {
		tx.Rollback()
	}
	defer stmt.Close()

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}
	defer stmt.Close()

	// Use a service account
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "para-pencari-jawaban"}

	sa := option.WithCredentialsFile("./database/user_service/para-pencari-jawaban-firebase-adminsdk-vavad-075df6c140.json")
	app, err := firebase.NewApp(ctx, conf, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	_, _, log_err := client.Collection("data").Add(ctx, map[string]interface{}{
		"createdAt": firestore.ServerTimestamp,
		"msg.log":   "add user",
	})
	if log_err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	defer stmt.Close()
	defer db.Close()

}

func GetUser(uuid int) model.UserNoPassword {
	user := model.UserNoPassword{}

	db := database.MySqlConnection()
	query := "select uuid, nik, nama, jenis_kelamin, alamat, role, saldo from users where uuid = ?"
	err := db.QueryRow(query, uuid).Scan(&user.Uuid, &user.Nik, &user.Nama, &user.Jenis_kelamin, &user.Alamat, &user.Role, &user.Saldo)
	if err != nil {
		panic(err)
	}

	return user
}
