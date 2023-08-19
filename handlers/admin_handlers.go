package handlers

import (
	"database/sql"
	"fmt"
	"golang_web_Project/auth"
	"golang_web_Project/database"
	"golang_web_Project/model"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func ShowUsersHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := database.MySqlConnection()
	var users []model.UserNoPassword
	tmpl, err := template.ParseFiles("templates/dashboardUser.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("CALL getAllUsers()")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var user model.UserNoPassword
		err := rows.Scan(&user.Nik, &user.Nama, &user.Jenis_kelamin, &user.Alamat, &user.Saldo)
		if err != nil {
			http.Redirect(w, r, "/error", http.StatusSeeOther)
			log.Fatal(err)
		}
		users = append(users, user)
	}
	fmt.Println(users)

	err = tmpl.Execute(w, users)
}

func ShowAddUserFormHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	baseTmpl, err := template.ParseFiles("templates/tambahPengguna.html")
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	err = baseTmpl.Execute(w, nil)
}

func SubmitAddUserFormHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := database.MySqlConnection()

	var userData model.User
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hashedPassword, error := auth.HashPassword(r.PostForm.Get("nik"))
	if err != nil {
		panic(error)
	}

	userData.Nama = r.PostForm.Get("nama")
	userData.Nik = r.PostForm.Get("nik")
	userData.Jenis_kelamin = r.PostForm.Get("gender")
	userData.Alamat = r.PostForm.Get("alamat")
	userData.Password = hashedPassword
	fmt.Println(userData)
	_, err = db.Exec("CALL InsertUser(?, ?, ?, ?, ?)", userData.Nik, userData.Nama, userData.Password, userData.Jenis_kelamin, userData.Alamat)
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		log.Fatal(err)
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func ShowUpdateUserFormHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := database.MySqlConnection()
	var user model.UserNoPassword
	queryParameters := r.URL.Query()
	paramName := queryParameters.Get("nik")
	err := r.ParseForm()

	tmpl, err := template.ParseFiles("templates/editPengguna.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	row := db.QueryRow("CALL getUserData(?)", paramName)
	row.Scan(&user.Nik, &user.Nama, &user.Jenis_kelamin, &user.Alamat, &user.Saldo)
	err = tmpl.Execute(w, user)
}

func SubmitUpdateUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
}
func DeleteUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := database.MySqlConnection()
	err := r.ParseForm()
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
	}
	nik := r.PostForm.Get("nik")

	_, deleteErr := db.Exec("CALL deleteUser(?)", nik)
	if deleteErr != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
	}
	http.Redirect(w, r, "/dashboard-user", http.StatusSeeOther)

}
func ShowTransactionsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := database.MySqlConnection()
	var rawTransactionDate []byte
	var transactions []model.Transaction

	tmpl, err := template.ParseFiles("templates/dashboardSaldo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rows, err := db.Query("CALL getAllTransaction()")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var transaction model.Transaction
		err := rows.Scan(&transaction.Nik, &transaction.Nama, &transaction.Jenis_transaksi, &transaction.Jumlah, &rawTransactionDate)
		if err != nil {
			http.Redirect(w, r, "/error", http.StatusSeeOther)
			log.Fatal(err)
		}
		transactionDateStr := string(rawTransactionDate)
		transaction.Tanggal, err = time.Parse("2006-01-02 15:04:05", transactionDateStr)
		transactions = append(transactions, transaction)
	}
	err = tmpl.Execute(w, transactions)
}

func ShowAddDepositHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := database.MySqlConnection()
	type User struct {
		Nik  string
		Nama string
	}

	var users []User

	rows, err := db.Query("CALL getUserDataForDeposit()")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Nik, &user.Nama)
		if err != nil {
			http.Redirect(w, r, "/error", http.StatusSeeOther)
			log.Fatal(err)
		}
		users = append(users, user)
	}
	tmpl, err := template.ParseFiles("templates/addSaldo.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, users)
}

func SubmitDepositHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := database.MySqlConnection()

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		return
	}

	nik := r.Form.Get("nik")
	jumlah := r.Form.Get("saldo")
	berat := r.Form.Get("berat")
	fmt.Println(nik)
	fmt.Println(jumlah)
	fmt.Println(berat)
	_, depositErr := db.Exec("CALL Deposit(?, ?, ?)", nik, jumlah, berat)
	if depositErr != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		log.Fatal(err)
	}

	http.Redirect(w, r, "/dashboard-transaksi", http.StatusSeeOther)

}

func ShowAddWithdrawFormHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := database.MySqlConnection()
	type User struct {
		Nik  string
		Nama string
	}

	var users []User

	rows, err := db.Query("CALL getUserDataForDeposit()")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Nik, &user.Nama)
		if err != nil {
			http.Redirect(w, r, "/error", http.StatusSeeOther)
			log.Fatal(err)
		}
		users = append(users, user)
	}
	tmpl, err := template.ParseFiles("templates/addWithdraw.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, users)

}

func SubmitWithdrawHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := database.MySqlConnection()
	var saldo int
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		return
	}

	nik := r.Form.Get("nik")
	jumlahstr := r.Form.Get("jumlah")
	jumlah, err := strconv.Atoi(jumlahstr)

	err = db.QueryRow("CALL getUserSaldoForWithdraw(?)", nik).Scan(&saldo)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Redirect(w, r, "/error", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/error", http.StatusSeeOther)
		}
	}
	if saldo >= jumlah {
		_, err := db.Exec("CALL Withdraw(?, ?)", nik, jumlah)
		if err != nil {
			http.Redirect(w, r, "/error", http.StatusSeeOther)
		}
	} else {
		fmt.Println("saldo tidak cukup")
		http.Redirect(w, r, "/dashboard-transaksi", http.StatusSeeOther)
	}
	http.Redirect(w, r, "/dashboard-transaksi", http.StatusSeeOther)
}
