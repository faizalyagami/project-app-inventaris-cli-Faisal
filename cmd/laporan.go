package cmd

import (
	"fmt"
	"project-app-inventaris-cli-faisal/handler"

	"github.com/spf13/cobra"
)

var laporanCmd = &cobra.Command{
	Use:   "laporan",
	Short: "Menu laporan & pengecekan barang",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			fmt.Println("--- Menu Laporan & Pengecekan ---")
			fmt.Println("1. Barang yang perlu diganti (>100 hari)")
			fmt.Println("2. Total nilai investasi & depresiasi semua barang")
			fmt.Println("3. Laporan investasi & depresiasi berdasarkan ID")
			fmt.Println("0. Kembali")

			var choice string
			fmt.Print("Pilih menu laporan: ")
			fmt.Scanln(&choice)

			switch choice {
			case "1":
				handler.LaporanBarangDiganti()
			case "2":
				handler.LaporanTotalInvestasi()
			case "3":
				handler.LaporanPerBarang()
			case "0":
				return
			default:
				fmt.Println("Pilihan tidak valid.")
			}
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(laporanCmd)
}
