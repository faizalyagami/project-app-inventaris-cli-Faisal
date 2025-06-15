package service

import (
	"bufio"
	"fmt"
	"os"
	"project-app-inventaris-cli-faisal/model"
	"project-app-inventaris-cli-faisal/repository"
	"strconv"
	"strings"
)

func ListBarang() {
	barang, err := repository.GetAllBarang()
	if err != nil {
		fmt.Println("Gagal mengambil data: ", err)
		return
	}
	fmt.Println("Daftar Barang:")
	for _, b := range barang {
		fmt.Printf("- %s (%s) | Kategori ID: %d\n", b.NamaBarang, b.KodeBarang, b.IDKategori)
	}
}

func AddBarang()  {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nama Barang: ")
	nama, _ := reader.ReadString('\n')

	fmt.Print("Kode Barang: ")
	kode, _ := reader.ReadString('\n')

	fmt.Print("ID Kategori: ")
	kategoriStr, _ := reader.ReadString('\n')
	idKategori, _ := strconv.Atoi(strings.TrimSpace(kategoriStr))

	data := model.Barang{
		NamaBarang: strings.TrimSpace(nama),
		KodeBarang: strings.TrimSpace(kode),
		IDKategori: idKategori,
	}

	if err := repository.InsertBarang(data); err != nil {
		fmt.Println("Gagal menambah barang:", err)
	} else {
		fmt.Println("Barang berhasil ditambahkan.")
	}
}

func DetailBarang()  {
	var id int
	fmt.Print("Masukkan ID barang: ")
	fmt.Scan(&id)
	b, err := repository.GetBarangByID(id)
	if err != nil {
		fmt.Println("Gagal mengambil data:", err)
	} else {
		fmt.Printf("Nama Barang: %s\nKode: %s\nID Kategori: %d\n", b.NamaBarang, b.KodeBarang, b.IDKategori)
	}

	// fmt.Println("\nTekan ENTER untuk kembali ke menu...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func UpdateBarang()  {
	var id int
	fmt.Print("Masukkan ID barang yang akan diupdate: ")
	fmt.Scan(&id)

	fmt.Scanln()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nama Baru: ")
	nama, _ :=reader.ReadString('\n')

	fmt.Print("Kode Baru: ")
	kode, _ :=reader.ReadString('\n')

	fmt.Print("ID Kategori Baru: ")
	kategoriStr, _ := reader.ReadString('\n')
	idKategori, _ := strconv.Atoi(strings.TrimSpace(kategoriStr))

	if !repository.KategoriExist(idKategori) {
		fmt.Println("ID Kategori tidak ditemukan.")
		return
	}

	b := model.Barang{
		ID: id,
		NamaBarang: strings.TrimSpace(nama),
		KodeBarang: strings.TrimSpace(kode),
		IDKategori: idKategori,
	}

	if err := repository.UpdateBarang(b); err != nil {
		fmt.Println("Update gagal: ", err)
	} else {
		fmt.Println("Update berhasil.")
	}
}

func DeleteBarang()  {
	var id int
	fmt.Print("Masukkan ID barang yang akan dihapus:")
	fmt.Scan(&id)

	fmt.Scanln()

	if err := repository.DeleteBarang(id); err != nil {
		fmt.Println("Gagal menghapus barang:", err)
	} else {
		fmt.Println("Barang berhasil dihapus.")
	}
}