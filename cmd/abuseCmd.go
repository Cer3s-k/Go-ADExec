package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var abuseCmd = &cobra.Command{
	Use:   "abuse",
	Short: "abuse the ActiveDirectory configuration",
	Long:  "abuse the ActiveDirectory configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[+] Go-ADExec abuse -h")
	},
}
