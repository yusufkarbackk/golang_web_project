package model

import (
// "database/sql"
// "time"
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
	Uuid          int
	Nik           string
	Nama          string
	Role          string
	Saldo         int
	Jenis_kelamin string
	Alamat        string
}
