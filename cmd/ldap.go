package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ldapCmd = &cobra.Command{
	Use:   "ldap",
	Short: "ldap query ActiveDirectory configuration",
	Long:  "ldap query ActiveDirectory configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[+] Go-ADExec ldap -h")
	},
}
