package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var abuseCmd = &cobra.Command{
	Use:   "abuse",
	Short: "redTeam abuse",
	Long:  "abuse the ActiveDirectory configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[+] search.exe -h")
	},
}
