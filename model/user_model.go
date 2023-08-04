package model

import (
// "database/sql"
// "time"
)

type User struct {
	Uuid     int
	Nik      string
	Nama     string
	Password string
	Role     string
	Saldo    int
}
