package cmd

import (
	"Go-ADExec/colors"
	"github.com/spf13/cobra"
	"os"
)

// smb module
var smbCmd = &cobra.Command{
	Use:   "smb",
	Short: "smb shares related attacks",
	Long:  "smb shares related attacks",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := parseGlobalInfo(cmd)
		if err != nil {
			colors.ErrorPrintln("smb initialization failed...")
			colors.ErrorPrintln("Go-ADExec smb -h")
			os.Exit(1)
		}
		if config == nil {
			colors.ErrorPrintln("config nil")
		}

		err = LdapConnect(config)
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {

	//store common parameters into a structure
	smbCmd.Flags().StringVarP(&GlobalLoginInfo.Domain, "domain", "d", "", "domain name")
	smbCmd.Flags().StringVarP(&GlobalLoginInfo.User, "username", "u", "", "domain username")
	smbCmd.Flags().StringVarP(&GlobalLoginInfo.Pass, "password", "p", "", "use user password")
	smbCmd.Flags().StringVarP(&GlobalLoginInfo.Hash, "hashes", "H", "", "use user hashes")
	smbCmd.Flags().StringVarP(&GlobalLoginInfo.GssApi, "gssapi", "g", "", "specified domain controller for authenticated with the current user's credentials.(example: --gssapi dc.test.local)")
	smbCmd.Flags().BoolVarP(&GlobalLoginInfo.SSLCon, "ssl", "s", false, "Use ssl to connect to ldap. default false")
	smbCmd.Flags().StringVarP(&GlobalLoginInfo.BaseDN, "basedn", "b", "", "Specify DN (ou=xx,dc=xx,dc=xx)")
	smbCmd.Flags().StringVarP(&GlobalLoginInfo.Output, "output", "o", "", "save result to file.")

}
