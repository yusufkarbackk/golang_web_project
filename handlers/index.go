package handlers

import (
	"fmt"
	"golang_web_Project/auth"
	"golang_web_Project/database"
	"golang_web_Project/model"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := database.MySqlConnection()
	var rawTransactionDate []byte
	var userTransactionData []model.Data
	type UserData struct {
		TransactionData []model.Data
		Nama            string
		Saldo           int
	}
	session, err := auth.Store.Get(r, "login_session")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	rows, err := db.Query("CALL getTransactions(?)", session.Values["nik"].(string))
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var data model.Data
		err := rows.Scan(&data.TransactionId, &data.TransactionType, &data.Berat, &data.Amount, &rawTransactionDate)
		if err != nil {
			log.Fatal(err)
		}
		transactionDateStr := string(rawTransactionDate)
		data.TransactionDate, err = time.Parse("2006-01-02 15:04:05", transactionDateStr)
		userTransactionData = append(userTransactionData, data)
	}
	fmt.Println(userTransactionData)
	userData := UserData{
		TransactionData: userTransactionData,
		Nama:            session.Values["nama"].(string),
		Saldo:           session.Values["saldo"].(int),
	}
	tmpl, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println(userData)

	err = tmpl.Execute(w, userData)
}
