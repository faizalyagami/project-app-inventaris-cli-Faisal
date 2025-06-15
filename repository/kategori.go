package repository

import (
	"project-app-inventaris-cli-faisal/database"
	"project-app-inventaris-cli-faisal/model"
)

func GetAllKategori() ([]model.Kategori, error) {
	db := database.GetDb()
	rows, err := db.Query("SELECT id, nama FROM kategori")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kategoris []model.Kategori
	for rows.Next() {
		var k model.Kategori
		if err := rows.Scan(&k.ID, &k.Nama); err != nil {
			return nil, err
		}
		kategoris = append(kategoris, k)
	}
	return kategoris, nil
}

func CreateKategori(nama string) error {
	db := database.GetDb()
	_, err := db.Exec("INSERT INTO kategori (nama) VALUES ($1)", nama)
	return err
}

func GetKategoriByID(id int) (model.Kategori, error) {
	db := database.GetDb()
	var k model.Kategori
	err := db.QueryRow("SELECT id, nama FROM kategori WHERE id = $1", id).Scan(&k.ID, &k.Nama)
	return k, err
}

func UpdateKategori(id int, nama string) error {
	db := database.GetDb()
	_,err := db.Exec("UPDATE kategori SET nama = $1 WHERE id = $2", nama, id)
	return err
}

func KategoriExist(id int) bool {
	db := database.GetDb()
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM kategori WHERE id=$1)", id).Scan(&exists)
	return err == nil && exists
}

func DeleteKategori(id int) error {
	db := database.GetDb()
	_, err := db.Exec("DELETE FROM kategori WHERE id = $1", id)
	return err
}