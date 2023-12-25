package cmd

import (
	"Go-ADExec/colors"
	"github.com/spf13/cobra"
)

var ldapCmd = &cobra.Command{
	Use:   "ldap",
	Short: "ldap query ActiveDirectory configuration",
	Long:  "ldap query ActiveDirectory configuration",
	Run: func(cmd *cobra.Command, args []string) {
		colors.PrintErrorf("Go-ADExec ldap -h")
	},
}
