package service

import (
	"project-app-inventaris-cli-faisal/model"
	"project-app-inventaris-cli-faisal/repository"
)

func GetAllKategori() ([]model.Kategori, error) {
	return repository.GetAllKategori()
}

func CreateKategori(nama string) error {
	return repository.CreateKategori(nama)
}

func GetKategoriByID(id int) (model.Kategori, error) {
	return repository.GetKategoriByID(id)
}

func UpdateKategori(id int, nama string) error {
	return repository.UpdateKategori(id, nama)
}

func DeleteKategori(id int) error {
	return repository.DeleteKategori(id)
}