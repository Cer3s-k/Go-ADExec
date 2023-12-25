package cmd

import (
	"Go-ADExec/colors"
	"github.com/spf13/cobra"
)

var adcsCmd = &cobra.Command{
	Use:   "adcs",
	Short: "Certificate related attacks",
	Long:  "Certificate related attacks",
	Run: func(cmd *cobra.Command, args []string) {
		colors.PrintErrorf("Go-ADExec adcs -h")
	},
}
