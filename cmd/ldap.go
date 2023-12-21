package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ldapCmd = &cobra.Command{
	Use:   "ldap",
	Short: "ldap query",
	Long:  "ldap query ActiveDirectory configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[+] ldap.exe -h")
	},
}
