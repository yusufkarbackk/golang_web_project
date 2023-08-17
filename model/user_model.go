package model

import (
	// "database/sql"
	"time"
)

type User struct {
	Uuid          int
	Nik           string
	Nama          string
	Jenis_kelamin string
	Alamat        string
	Password      string
	Role          string
	Saldo         int
}

type UserNoPassword struct {
	Nik           string
	Nama          string
	Jenis_kelamin string
	Alamat        string
	Saldo         int
}

type Data struct {
	TransactionId   int
	TransactionType string
	Berat           int
	Amount          int
	TransactionDate time.Time
}
