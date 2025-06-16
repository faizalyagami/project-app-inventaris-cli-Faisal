package repository_test

import (
	"project-app-inventaris-cli-faisal/database"
	"project-app-inventaris-cli-faisal/repository"
	"testing"
)

func TestKategoriExist(t *testing.T) {
	database.InitDb()
	exists := repository.KategoriExist(1)
	if !exists {
		t.Errorf("Expected kategori ID 1 to exist")
	}
}