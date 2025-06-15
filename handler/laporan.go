package handler

import (
	"fmt"
	"project-app-inventaris-cli-faisal/repository"
	"project-app-inventaris-cli-faisal/service"
)

func LaporanBarangDiganti() {
	fmt.Println(">> Barang yang perlu diganti (>100 hari):")
	service.TampilkanBarangPerluDiganti()
	promptEnter()
}

func LaporanTotalInvestasi() {
	fmt.Println(">> Total nilai investasi dan depresiasi:")
	service.TampilkanTotalInvestasi()
	promptEnter()
}

func LaporanPerBarang() {
	var id string
	fmt.Print("Masukkan ID barang: ")
	fmt.Scanln(&id)
	service.TampilkanLaporanPerBarang(id)
	promptEnter()
}

func LaporanBarangLebih100Hari()  {
	fmt.Println(">> Menampilkan barang yang sudah >100 har")

	barangs := repository.GetBarangPerluDiganti()
	if len(barangs) == 0 {
		fmt.Println("Tidak ada barang yang sudah >100 hari.")
	} else {
		fmt.Println("ID\tNama\t\tUmur (hari)")
		for _, b := range barangs {
			fmt.Printf("%s\t%s\t\t%d\n", b.ID, b.Nama, b.UmurHari)
		}
	}
	promptEnter()
}

func promptEnter() {
	fmt.Print("\nTekan ENTER untuk kembali ke menu...")
	fmt.Scanln()
}