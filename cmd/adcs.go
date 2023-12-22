package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var adcsCmd = &cobra.Command{
	Use:   "adcs",
	Short: "Certificate related attacks",
	Long:  "Certificate related attacks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[+] Go-ADExec adcs -h")
	},
}
