package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var smbCmd = &cobra.Command{
	Use:   "smb",
	Short: "smb shares related attacks",
	Long:  "smb shares related attacks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[+] Go-ADExec smb -h")
	},
}
