package repository

import (
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