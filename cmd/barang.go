package cmd

import (
	"fmt"
	"project-app-inventaris-cli-faisal/handler"

	"github.com/spf13/cobra"
)

var barangCmd = &cobra.Command{
	Use: "barang",
	Short: "Kelola data barang inventaris",
	Run: func (cmd *cobra.Command, args []string)  {
		for {
			fmt.Println("--- Menu Barang Inventaris ---")
			fmt.Println("1. Lihat semua barang")
			fmt.Println("2. Tambah barang")
			fmt.Println("3. Lihat detail barang")
			fmt.Println("4. Update barang")
			fmt.Println("5. Hapus barang")
			fmt.Println("0. Kembali")

			var choice string
			fmt.Print("Pilih menu barang: ")
			fmt.Scanln(&choice)

			switch choice {
			case "1":
				handler.ListBarang()
			case "2":
				handler.AddBarang()
			case "3":
				handler.DetailBarang()
			case "4":
				handler.UpdateBarang()
			case "5":
				handler.DeleteBarang()
			case "0":
				return
			default:
				fmt.Println("Pilihan tidak valid.")
			}
			fmt.Println()
		}
	},
}

func init()  {
	rootCmd.AddCommand(barangCmd)
}