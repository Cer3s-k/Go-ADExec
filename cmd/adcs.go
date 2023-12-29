package cmd

import (
	"Go-ADExec/colors"
	"github.com/spf13/cobra"
)

// adcs module
var adcsCmd = &cobra.Command{
	Use:   "adcs",
	Short: "Certificate related attacks",
	Long:  "Certificate related attacks",
	Run: func(cmd *cobra.Command, args []string) {
		colors.ErrorPrintln("Go-ADExec adcs -h")
	},
}
