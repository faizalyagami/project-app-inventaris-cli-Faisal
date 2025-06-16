package repository

import (
	"database/sql"
	"project-app-inventaris-cli-faisal/database"
	"project-app-inventaris-cli-faisal/model"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error saat setup mock DB: %s", err)
	}
	database.InitDb()
	return db, mock
}

func TestGetBarangByID(t *testing.T)  {
	db, mock := setupMockDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "nama_barang", "kode_barang", "id_kategori"}).AddRow(1, "Barang A", "BA001", 2)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, nama_barang, kode_barang, id_kategori FROM barang WHERE id=$1")).
		WithArgs(1).WillReturnRows(rows)
	
	result, err := GetBarangByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "Barang A", result.NamaBarang)
	assert.Equal(t, "BA001", result.KodeBarang)
	assert.Equal(t, 2, result.IDKategori)
}

func TestInsertBarang(t *testing.T)  {
	db, mock := setupMockDB(t)
	defer db.Close()

	b := model.Barang{NamaBarang: "Barang B", KodeBarang: "BB001", IDKategori: 1}

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO barang (nama_barang, kode_barang, id_kategori) VALUES ($1, $2, $3)")).
		WithArgs(b.NamaBarang, b.KodeBarang, b.IDKategori).WillReturnResult(sqlmock.NewResult(1,1))
	err := InsertBarang(b)
	assert.NoError(t, err)
}

func TestUpdateBarang(t *testing.T)  {
	db, mock := setupMockDB(t)
	defer db.Close()

	b :=model.Barang{ID: 1, NamaBarang: "Barang Update", KodeBarang: "BU001", IDKategori: 3}

	mock.ExpectExec(regexp.QuoteMeta("UPDATE barang SET nama_barang=$1, kode_barang=$2, id_kategori=$3 WHERE id=$4")).
		WithArgs(b.NamaBarang, b.KodeBarang, b.IDKategori, b.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	
	err := UpdateBarang(b)
	assert.NoError(t, err)
}

func TestDeleteByID(t *testing.T)  {
	db, mock := setupMockDB(t)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM barang WHERE id=$1")).WithArgs(1).WillReturnResult(sqlmock.NewResult(1,1))
	err := DeleteBarang(1)
	assert.NoError(t, err)
}

func TestCariBarangByNama(t *testing.T)  {
	db, mock := setupMockDB(t)
	defer db.Close()
	
	rows := sqlmock.NewRows([]string {
		"id", "nama_barang", "kode_barang", "nama_kategori", "umur_hari", "status",
	}).AddRow(1, "Barang C", "BC001", "Elektronik", 100, "Masih Layak")

	mock.ExpectQuery("SELECT .* FROM barang b JOIN kategori k .* WHERE LOWER\\(b\\.nama_barang\\) LIKE LOWER\\(\\$1\\)").WithArgs("%barang%").WillReturnRows(rows)
	
	results := CariBarangByNama("barang")

	assert.Len(t, results, 1)
	assert.Equal(t, "Barang C", results[0].NamaBarang)
	assert.Equal(t, "Masih Layak", results[0].Status)
}