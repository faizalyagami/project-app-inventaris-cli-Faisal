package handler

import (
	"bufio"
	"fmt"
	"os"
	"project-app-inventaris-cli-faisal/service"
	"strings"
)

func ListKategori() {
	kategoris, err := service.GetAllKategori()
	if err != nil {
		fmt.Println("Gagal mengambil data kategori:", err)
		return
	}
	fmt.Println("Daftar Kategori:")
	for _, k := range kategoris {
		fmt.Printf("- %d: %s\n", k.ID, k.NamaKategori)
	}
}

func AddKategori()  {
	fmt.Print("Masukkan Nama Kategori: ")
	reader := bufio.NewReader(os.Stdin)
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	err := service.CreateKategori(nama)
	if err != nil {
		fmt.Println("Gagal menambahkan kategori:", err)
		return
	}
	fmt.Println("Kategori berhasil ditambahkan")
}

func DetailKategori()  {
	id := readInt("Masukkan ID Kategori: ")
	kategori, err := service.GetKategoriByID(id)
	if err != nil {
		fmt.Println("Kategori tidak ditemukan:", err)
		return
	}
	fmt.Printf("ID: %d\nNama: %s\n", kategori.ID, kategori.NamaKategori)
}

func UpdateKategori()  {
	id := readInt("Masukkan ID kategori yang baru ingin diupdate: ")
	fmt.Print("Masukkan nama kategori baru: ")
	reader := bufio.NewReader(os.Stdin)
	nama, _ := reader.ReadString('\n')

	err := service.UpdateKategori(id, nama)
	if err != nil {
		fmt.Println("Gagal mengupdate kategori:", err)
		return
	}
	fmt.Println("Kategori berhasil diupdate.")
}

func DeleteKategori()  {
	id := readInt("Masukkan ID kategori yang ingin dihapus:")
	err := service.DeleteKategori(id)
	if err != nil {
		fmt.Println("Gagal menghapus kategori:", err)
		return
	}
	fmt.Println("Kategori berhasil dihapus")
}

func readInt(prompt string) int {
	fmt.Print(prompt)
	var input int
	fmt.Scanln(&input)
	return input
}