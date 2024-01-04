package cmd

import (
	"Go-ADExec/colors"
	"Go-ADExec/module/ldapaction"
	"github.com/spf13/cobra"
	"os"
)

// ldap module
var ldapCmd = &cobra.Command{
	Use:   "ldap",
	Short: "ldap query ActiveDirectory configuration",
	Long:  "ldap query ActiveDirectory configuration",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := ldapaction.ParseGlobalInfo(cmd)
		if err != nil {
			colors.ErrorPrintln(err)
			colors.ErrorPrintln("Go-ADExec ldap -h")
			os.Exit(1)
		}

		//initialize ldap connection
		err = ldapaction.LdapConnect(config)
		if err != nil {
			os.Exit(1)
		}

		//custom ldap search
		if ldapaction.QueryInfo.Attr.Attributes != nil && ldapaction.QueryInfo.Attr.Filter == "" {
			colors.ErrorPrintln("please enter filter criteria")
		} else if ldapaction.QueryInfo.Attr.Filter != "" {
			ldapaction.CustomSearch()
		}

	},
}

func init() {

	//store common parameters into a structure
	ldapCmd.PersistentFlags().StringVarP(&ldapaction.GlobalLoginInfo.Domain, "domain", "d", "", "domain name")
	ldapCmd.PersistentFlags().StringVarP(&ldapaction.GlobalLoginInfo.User, "username", "u", "", "domain username")
	ldapCmd.PersistentFlags().StringVarP(&ldapaction.GlobalLoginInfo.Pass, "password", "p", "", "use user password")
	ldapCmd.PersistentFlags().StringVarP(&ldapaction.GlobalLoginInfo.Hash, "hashes", "H", "", "use user hashes")
	ldapCmd.PersistentFlags().StringVarP(&ldapaction.GlobalLoginInfo.GssApi, "gssapi", "g", "", "specified domain controller for authenticated with the current user's credentials.(example: --gssapi dc.test.local)")
	ldapCmd.PersistentFlags().BoolVarP(&ldapaction.GlobalLoginInfo.SSLCon, "ssl", "s", false, "Use ssl to connect to ldapaction. default false")
	ldapCmd.PersistentFlags().StringVarP(&ldapaction.GlobalLoginInfo.BaseDN, "basedn", "b", "", "Specify DN (ou=xx,dc=xx,dc=xx)")
	ldapCmd.PersistentFlags().StringVarP(&ldapaction.GlobalLoginInfo.Output, "output", "o", "", "save result to file.")

	ldapCmd.Flags().StringVarP(&ldapaction.QueryInfo.Attr.Filter, "filter", "f", "", "use custom search query")
	ldapCmd.Flags().StringSliceVarP(&ldapaction.QueryInfo.Attr.Attributes, "attributes", "a", nil, "ldap attribute of search output(default: distinguishedName)")

	//common commands
	ldapCmd.AddCommand(ldapaction.DomainAdminCmd)
	ldapCmd.AddCommand(ldapaction.DomainControllersCmd)
	ldapCmd.AddCommand(ldapaction.DomainMAQCmd)
	ldapCmd.AddCommand(ldapaction.DomainOUCmd)
	ldapCmd.AddCommand(ldapaction.DomainComputerCmd)
	ldapCmd.AddCommand(ldapaction.DomainUserCmd)
	ldapCmd.AddCommand(ldapaction.DomainGroupCmd)
	ldapCmd.AddCommand(ldapaction.DomainBaseInfoCmd)
}
