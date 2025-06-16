package repository

import (
	"fmt"
	"project-app-inventaris-cli-faisal/database"
	"project-app-inventaris-cli-faisal/model"
)

type LaporanBarang struct {
	ID         string
	Nama       string
	UmurHari   int
	Investasi  float64
	Depresiasi float64
}

func GetBarangPerluDiganti() []LaporanBarang {
	return []LaporanBarang{
		{ID: "B001", Nama: "Laptop A", UmurHari: 120},
		{ID: "B002", Nama: "Printer B", UmurHari: 150},
	}
}

func HitungTotalInvestasiDanDepresiasi() (totalInvestasi, totalDepresiasi float64) {
	return 10000000, 2500000
}

func GetLaporanBarangByID(id int) model.LaporanBarang {
	// if id == "B001" {
	// 	return LaporanBarang{ID: "B001", Nama: "Laptop A", Investasi: 5000000, Depresiasi: 1000000}
	// }
	// return LaporanBarang{}
	var result model.LaporanBarang

	query := `
		SELECT
			b.id,
			b.nama_barang,
			b.harga AS investasi,
			ROUND((b.harga / b.umur_ekonomis) * 
			EXTRACT(DAY FROM NOW() - b.tanggal_beli), 2) AS despresiasi
		FROM barang b
		WHERE b.id = $1
	`

	row := database.GetDb().QueryRow(query, id)
	err := row.Scan(&result.ID, &result.NamaBarang, &result.Investasi, &result.Depresiasi)
	if err != nil {
		fmt.Println("Barang tidak ditemukan atau gagal diambil:", err)
		return model.LaporanBarang{}
	}
	return result
}
