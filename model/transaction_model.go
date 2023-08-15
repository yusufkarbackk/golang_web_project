package model

import "time"

type Transaction struct {
	Nik             int
	Nama            string
	Jenis_transaksi string
	Jumlah          int
	Tanggal         time.Time
}
