package model

import "time"

type Barang struct {
	ID          int
	NamaBarang  string
	KodeBarang  string
	IDKategori  int
	TanggalBeli time.Time
	Harga float64
	UmurEkonomis int
}
