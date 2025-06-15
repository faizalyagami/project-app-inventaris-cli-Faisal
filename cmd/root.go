package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "inventaris",
	Short: "Aplikasi CLI Inventaris Kantor",
}

func Execute()  {
	cobra.CheckErr(rootCmd.Execute())
}
