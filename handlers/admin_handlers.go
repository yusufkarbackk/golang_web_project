package handlers

import (
	"database/sql"
	"fmt"
	"golang_web_Project/auth"
	"golang_web_Project/database"
	// userservice "golang_web_Project/database/user_service"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	_, err = db.Exec("CALL InsertUser(?, ?, ?, ?, ?)", userData.Nik, userData.Nama, userData.Password, userData.Jenis_kelamin, userData.Alamat)
	if err != nil {
		log.Fatal(err)
		// redirect ke halaman error
	}
	http.Redirect(w, r, "/dashboard-user", http.StatusSeeOther)
}

func ShowUpdateUserFormHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// baseTmpl, err := template.ParseFiles("templates/editPengguna.html")
	id, err := strconv.Atoi(r.URL.Query().Get("uuid"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// data := userservice.GetUser(id)
	fmt.Println(id)
	// err = baseTmpl.Execute(w, data)
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
			fmt.Println(transaction)
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
		Nik  int
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
		log.Fatal(err)
		// redirect ke halaman error
	}
	
	http.Redirect(w, r, "/dashboard-transaksi", http.StatusSeeOther)

}

func ShowAddWithdrawFormHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := database.MySqlConnection()
	type User struct {
		Nik  int
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
	fmt.Println(jumlah)
	err = db.QueryRow("CALL getUserSaldoForWithdraw(?)", nik).Scan(&saldo)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found")
		} else {
			panic(err)
		}
	}
	fmt.Println(saldo)
	if saldo >= jumlah {
		_, err := db.Exec("CALL Withdraw(?, ?)", nik, jumlah)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("saldo tidak cukup")
		http.Redirect(w, r, "/dashboard-transaksi", http.StatusSeeOther)
	}
	http.Redirect(w, r, "/dashboard-transaksi", http.StatusSeeOther)
}
