package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"project-app-inventaris-cli-faisal/database"
	"project-app-inventaris-cli-faisal/handler"
	"runtime"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	database.InitDb()
	for {
		ClearScreen()
		showMainMenu()

		choice := readInput("Pilih menu (1-4), atau 0 untuk keluar: ")

		switch choice {
		case "1":
			handleKategoriMenu()
		case "2":
			handleBarangMenu()
		case "3":
			handleLaporanMenu()
		case "4":
			searchBarang()
		case "0":
			fmt.Println("Terima kasih telah menggunakan Aplikasi Inventaris CLI.")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
		// promptEnterToContinue()
	}
}

func showMainMenu() {
	fmt.Println("========================================")
	fmt.Println("     SISTEM INVENTARIS KANTOR CLI")
	fmt.Println("========================================")
	fmt.Println("1. Kelola Kategori Barang")
	fmt.Println("2. Kelola Barang Inventaris")
	fmt.Println("3. Laporan & Pengecekan")
	fmt.Println("4. Cari Barang")
	fmt.Println("0. Keluar")
	fmt.Println("========================================")
}

func handleKategoriMenu() {
	ClearScreen()
	fmt.Println("--- Menu Kategori Barang ---")
	fmt.Println("1. Lihat semua kategori")
	fmt.Println("2. Tambah kategori")
	fmt.Println("3. Lihat detail kategori")
	fmt.Println("4. Update kategori")
	fmt.Println("5. Hapus kategori")
	fmt.Println("0. Kembali")
	choice := readInput("Pilih menu kategori: ")

	switch choice {
	case "1":
		listKategori()
	case "2":
		addKategori()
	case "3":
		detailKategori()
	case "4":
		updateKategori()
	case "5":
		deleteKategori()
	case "0":
		return
	default:
		fmt.Println("Pilihan tidak valid.")
	}
	promptEnterToContinue()
}

func handleBarangMenu() {
	ClearScreen()
	fmt.Println("--- Menu Barang Inventaris ---")
	fmt.Println("1. Lihat semua barang")
	fmt.Println("2. Tambah barang")
	fmt.Println("3. Lihat detail barang")
	fmt.Println("4. Update barang")
	fmt.Println("5. Hapus barang")
	fmt.Println("0. Kembali")
	choice := readInput("Pilih menu barang: ")

	switch choice {
	case "1":
		listBarang()
	case "2":
		addBarang()
	case "3":
		detailBarang()
	case "4":
		updateBarang()
	case "5":
		deleteBarang()
	case "0":
		return
	default:
		fmt.Println("Pilihan tidak valid.")
	}
	promptEnterToContinue()
}

func handleLaporanMenu() {
	ClearScreen()
	fmt.Println("--- Menu Laporan & Pengecekan ---")
	fmt.Println("1. Barang yang perlu diganti (>100 hari)")
	fmt.Println("2. Total nilai investasi & depresiasi semua barang")
	fmt.Println("3. Laporan investasi & depresiasi berdasarkan ID")
	fmt.Println("0. Kembali")
	choice := readInput("Pilih menu laporan: ")

	switch choice {
	case "1":
		laporanBarangDiganti()
	case "2":
		laporanTotalInvestasi()
	case "3":
		laporanPerBarang()
	case "0":
		return
	default:
		fmt.Println("Pilihan tidak valid.")
	}
	promptEnterToContinue()
}

func readInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func promptEnterToContinue() {
	fmt.Print("\nTekan ENTER untuk kembali ke menu...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// --- Tambahan: ClearScreen untuk bersihkan layar terminal
func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// --- Placeholder Fungsi ---
func listKategori() {
	handler.ListKategori()
}
func addKategori() { 
	handler.AddKategori()
 }
func detailKategori() {
	handler.DetailKategori()
}
func updateKategori() {
	handler.UpdateKategori()
}
func deleteKategori() {
	handler.DeleteKategori()
}
func listBarang() {
	handler.ListBarang()
}
func addBarang() {
	handler.AddBarang()
}
func detailBarang() {
	handler.DetailBarang()
}
func updateBarang() {
	handler.UpdateBarang()
}
func deleteBarang() {
	handler.DeleteBarang()
}
func laporanBarangDiganti() {
	handler.LaporanBarangDiganti()
}
func laporanTotalInvestasi() {
	handler.LaporanTotalInvestasi()
}
func laporanPerBarang()      {
	handler.LaporanPerBarang()
}
func searchBarang(){ 
	fmt.Println(">> Cari barang berdasarkan nama")
	keyword := readInput("Cari barang: ")

	handler.SearchBarangByNama(keyword)
	promptEnterToContinue()
}