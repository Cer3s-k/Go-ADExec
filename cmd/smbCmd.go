package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

//var userNameStr = "username"

func init() {
	smbCmd.Flags().StringVarP(&GlobalLoginInfo.DomainName, "domain", "d", "", "domain name")
	smbCmd.Flags().StringVarP(&GlobalLoginInfo.UserName, "username", "u", "", "domain username")
	smbCmd.Flags().StringVarP(&GlobalLoginInfo.UserPass, "password", "p", "", "use user password")
	smbCmd.Flags().StringVarP(&GlobalLoginInfo.UserHash, "hashes", "n", "", "use user ntlm hashes")
	smbCmd.Flags().StringVarP(&GlobalLoginInfo.GssApi, "gssapi", "g", "", "specified domain controller for authenticated with the current user's credentials.(example: --gssapi dc.test.local)")
	smbCmd.Flags().BoolVarP(&GlobalLoginInfo.SSLCon, "ssl", "s", false, "Use ssl to connect to ldap. default false")
	smbCmd.Flags().StringVarP(&GlobalLoginInfo.BaseDN, "basedn", "b", "", "Specify DN (ou=xx,dc=xx,dc=xx)")
	smbCmd.Flags().StringVarP(&GlobalLoginInfo.Output, "output", "o", "", "save result to file.")
}

var smbCmd = &cobra.Command{
	Use:   "smb",
	Short: "smb shares related attacks",
	Long:  "smb shares related attacks",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := parseGlobalInfo(cmd)
		if err != nil {
			fmt.Println("[-] smb initialization failed...")
			fmt.Println("[+] Go-ADExec smb --help")
			os.Exit(1)
		}
		if config == nil {
			fmt.Println("config nil")
		}
		_, err = LdapConnect(config)
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(config.UserName)
		fmt.Println(config.UserPass)
	},
}
