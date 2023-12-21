package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var smbCmd = &cobra.Command{
	Use:   "smb",
	Short: "smb detection",
	Long:  "attack using smb shares",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[+] search.exe -h")
	},
}
