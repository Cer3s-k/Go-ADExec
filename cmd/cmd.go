package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {

	//ldap module
	rootCmd.AddCommand(ldapCmd)

	// abuse module
	rootCmd.AddCommand(abuseCmd)

	//smb module
	rootCmd.AddCommand(smbCmd)

	// Disable the default generated completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	//rootCmd.Flags().

}

var rootCmd = &cobra.Command{
	Use:     "Go-ADExec",
	Short:   "Golang AD tools",
	Long:    "Go-ADExec是一款用Go编写的内网信息收集和利用工具",
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[+] Go-ADExec.exe -h")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
