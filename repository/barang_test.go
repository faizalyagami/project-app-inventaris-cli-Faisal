package repository_test

import (
	"project-app-inventaris-cli-faisal/database"
	"project-app-inventaris-cli-faisal/repository"
	"testing"
)

func TestGetAllBarang(t *testing.T) {
	database.InitDb()
	barang, err := repository.GetAllBarang()
	if err != nil {
		t.Errorf("Gagal mengambil data barang: %v", err)
	}

	if len(barang) == 0 {
		t.Log("Tidak ada barang ditemukan.")
	}
}