package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

func init() {

	//add ldap、abuse、smb、adcs module
	rootCmd.AddCommand(ldapCmd)
	rootCmd.AddCommand(ntlmCmd)
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
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
