package cmd

import (
	"Go-ADExec/module/ldapaction"
	"github.com/spf13/cobra"
)

// smb module
var smbCmd = &cobra.Command{
	Use:   "smb",
	Short: "smb shares related attacks",
	Long:  "smb shares related attacks",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {

	//store common parameters into a structure
	smbCmd.PersistentFlags().StringVarP(&ldapaction.GlobalLoginInfo.Domain, "domain", "d", "", "domain name")
	smbCmd.PersistentFlags().StringVarP(&ldapaction.GlobalLoginInfo.User, "username", "u", "", "domain username")
	smbCmd.PersistentFlags().StringVarP(&ldapaction.GlobalLoginInfo.Pass, "password", "p", "", "use user password")
	smbCmd.PersistentFlags().StringVarP(&ldapaction.GlobalLoginInfo.Hash, "hashes", "H", "", "use user hashes")
	smbCmd.PersistentFlags().StringVarP(&ldapaction.GlobalLoginInfo.GssApi, "gssapi", "g", "", "specified domain controller for authenticated with the current user's credentials.(example: --gssapi dc.test.local)")
	smbCmd.PersistentFlags().BoolVarP(&ldapaction.GlobalLoginInfo.SSLCon, "ssl", "s", false, "Use ssl to connect to ldapaction. default false")
	smbCmd.PersistentFlags().StringVarP(&ldapaction.GlobalLoginInfo.BaseDN, "basedn", "b", "", "Specify DN (ou=xx,dc=xx,dc=xx)")
	smbCmd.PersistentFlags().StringVarP(&ldapaction.GlobalLoginInfo.Output, "output", "o", "", "save result to file.")

}
