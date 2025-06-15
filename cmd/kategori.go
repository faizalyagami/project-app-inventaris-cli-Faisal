package cmd

import (
	"fmt"
	"project-app-inventaris-cli-faisal/handler"

	"github.com/spf13/cobra"
)

var kategoriCmd = &cobra.Command{
	Use: "kategori",
	Short: "Kelola kategori barang",
	Run: func (cmd *cobra.Command, args []string)  {
		for {
			fmt.Println("--- Menu Kategori Inventaris ---")
			fmt.Println("1. Lihat semua kategori")
			fmt.Println("2. Tambah kategori")
			fmt.Println("3. Lihat detail kategori")
			fmt.Println("4. Update kategori")
			fmt.Println("5. Hapus kategori")
			fmt.Println("0. Kembali")

			var choice string
			fmt.Print("Pilih menu kategori: ")
			fmt.Scanln(&choice)

			switch choice {
			case "1":
				handler.ListKategori()
			case "2":
				handler.AddKategori()
			case "3":
				handler.DeleteKategori()
			case "4":
				handler.UpdateKategori()
			case "5":
				handler.DeleteKategori()
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
	rootCmd.AddCommand(kategoriCmd)
}