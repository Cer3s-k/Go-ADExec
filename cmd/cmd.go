package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {

	//add ldap、abuse、smb、adcs module
	rootCmd.AddCommand(ldapCmd)
	rootCmd.AddCommand(abuseCmd)
	rootCmd.AddCommand(smbCmd)
	rootCmd.AddCommand(adcsCmd)

	// Disable the default generated completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true

}

var rootCmd = &cobra.Command{
	Use:     "Go-ADExec",
	Short:   "Golang AD tools",
	Long:    "Go-ADExec是一款用Go编写的内网信息收集和利用工具",
	Version: "1.0.0",
	//ValidArgs: []string{"abuse", "adcs", "ldap", "smb"},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
