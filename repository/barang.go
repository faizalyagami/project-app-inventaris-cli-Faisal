package repository

import (
	"fmt"
	"project-app-inventaris-cli-faisal/database"
	"project-app-inventaris-cli-faisal/model"
)

func GetAllBarang() ([]model.Barang, error) {
	db := database.GetDb()
	rows, err := db.Query("SELECT id, nama_barang, kode_barang, id_kategori FROM barang")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var barangList []model.Barang
	for rows.Next() {
		var b model.Barang
		err := rows.Scan(&b.ID, &b.NamaBarang, &b.KodeBarang, &b.IDKategori)
		if err != nil {
			return nil, err
		}
		barangList = append(barangList, b)
	}
	return barangList, nil
 }

 func GetBarangByID(id int) (model.Barang, error) {
	db := database.GetDb()
	var b model.Barang
	err := db.QueryRow("SELECT id, nama_barang, kode_barang, id_kategori FROM barang WHERE id=$1", id).
		Scan(&b.ID, &b.NamaBarang, &b.KodeBarang, &b.IDKategori)
	if err != nil {
		return b, err
	}
	return b, nil
 }

 func InsertBarang(b model.Barang) error {
	db := database.GetDb()
	_, err := db.Exec("INSERT INTO barang (nama_barang, kode_barang, id_kategori) VALUES ($1, $2, $3)", b.NamaBarang, b.KodeBarang)
	return err
 }

 func UpdateBarang(b model.Barang) error {
	db := database.GetDb()
	_, err := db.Exec("UPDATE barang SET nama_barang=$1, kode_barang=$2, id_kategori=$3 WHERE id=$4", b.NamaBarang, b.KodeBarang, b.IDKategori, b.ID)
	return err
 }

 func DeleteBarang(id int) error {
	db := database.GetDb()
	_, err := db.Exec("DELETE FROM barang WHERE id=$1", id)
	return err
 }

 func CariBarangByNama(keyword string) []model.BarangWithKategori {
	var results []model.BarangWithKategori
	query := 
	`
	SELECT 
		b.id,
		b.nama_barang,
		b.kode_barang,
		k.nama_kategori,
		EXTRACT(DAY FROM NOW() - b.tanggal_beli) AS umur_hari,
		CASE
			WHEN EXTRACT(DAY FROM NOW() - b.tanggal_beli) > b.umur_ekonomis THEN 'Perlu Diganti'
			ELSE 'Masih Layak'
		END AS status
	FROM barang b
	JOIN kategori k ON b.id_kategori = k.id
	WHERE LOWER(b.nama_barang) LIKE LOWER($1)

	`
	rows, err := database.GetDb().Query(query, "%"+keyword+"%")
	if err != nil {
		fmt.Println("Gagal melakukan pencairan:", err)
		return results
	}
	defer rows.Close()

	for rows.Next() {
		var b model.BarangWithKategori
		err := rows.Scan(&b.ID, &b.NamaBarang, &b.KodeBarang, &b.NamaKategori, &b.UmurHari, &b.Status)
		if err != nil {
			fmt.Println("Gagal membaca hasil baris:", err)
			continue
		}
		results = append(results, b)
	}
	return results
 }