package ldapaction

import (
	"Go-ADExec/colors"
	"github.com/spf13/cobra"
	"os"
)

type SearchAttr struct {
	// filter
	Filter string
	// attribute
	Attributes []string
}

type queryConfig struct {
	Global *LdapInfo
	Attr   *SearchAttr
}

// QueryInfo global variables for ldap search
var QueryInfo = queryConfig{
	Global: &GlobalLoginInfo,
	Attr:   &SearchAttr{Attributes: []string{"distinguishedName"}},
}

var UserInfoCmd = &cobra.Command{
	Use:   "DA",
	Short: "查询域内域管组用户信息",
	Long:  "查询域内域管组用户信息",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := ParseGlobalInfo(cmd)
		if err != nil {
			colors.ErrorPrintln(err)
			colors.ErrorPrintln("Go-ADExec ldap user -h")
			os.Exit(1)
		}

		//initialize ldap connection
		err = LdapConnect(config)
		if err != nil {
			os.Exit(1)
		}

		domainAdminQuery(&QueryInfo)
	},
}
