package service

import (
	"fmt"
	"os"
	"project-app-inventaris-cli-faisal/repository"
	"strconv"
	"text/tabwriter"
)

func TampilkanBarangPerluDiganti() {
	data := repository.GetBarangPerluDiganti()
	if len(data) == 0 {
		fmt.Println("Tidak ada barang yang perlu diganti (>100 hari).")
		return
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(writer, "ID\tNama Barang\tUmur (Hari)")
	fmt.Fprintln(writer, "--\t------------\t------------")

	for _, b := range data {
		fmt.Fprintf(writer, "%s\t%s\t%d\n", b.ID, b.Nama, b.UmurHari)
	}
	writer.Flush()
}

func TampilkanTotalInvestasi() {
	total, depresiasi := repository.HitungTotalInvestasiDanDepresiasi()

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(writer, "Total Investasi\tTotal Depresiasi")
	fmt.Fprintln(writer, "---------------\t-----------------")
	fmt.Fprintf(writer, "Rp %.2f\tRp %.2f\n", total, depresiasi)
	writer.Flush()
}

func TampilkanLaporanPerBarang(idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID tidak valid.")
		return
	}

	laporan := repository.GetLaporanBarangByID(id)
	if laporan.ID == 0 {
		fmt.Println("Barang tidak ditemukan.")
		return
	}
	fmt.Printf("ID: %d\nNama: %s\nInvestasi: Rp %.2f\nDepresiasi: Rp %.2f\n", laporan.ID, laporan.NamaBarang, laporan.Investasi, laporan.Depresiasi)
	// laporan := repository.GetLaporanBarangByID(id)
	// if laporan.ID == "" {
	// 	fmt.Println("Barang tidak ditemukan.")
	// 	return
	// }

	// writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	// fmt.Fprintln(writer, "ID\tNama\tInvestasi\tDepresiasi")
	// fmt.Fprintln(writer, "--\t----\t----------\t-----------")
	// fmt.Fprintf(writer, "%s\t%s\tRp %.2f\tRp %.2f\n", laporan.ID, laporan.Nama, laporan.Investasi, laporan.Depresiasi)
	// writer.Flush()
}

func HitungTotalInvestasiDanDepresiasi()  {
	
}