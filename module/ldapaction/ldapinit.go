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

// DomainAdminCmd query domain admin
var DomainAdminCmd = &cobra.Command{
	Use:   "DA",
	Short: "查询域内域管组用户",
	Long:  "查询域内域管组用户",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := ParseGlobalInfo(cmd)
		if err != nil {
			colors.ErrorPrintln(err)
			colors.ErrorPrintln("Go-ADExec ldap DA -h")
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

// DomainControllersCmd query domain controllers
var DomainControllersCmd = &cobra.Command{
	Use:   "DC",
	Short: "查询域控",
	Long:  "查询域控",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := ParseGlobalInfo(cmd)
		if err != nil {
			colors.ErrorPrintln(err)
			colors.ErrorPrintln("Go-ADExec ldap DC -h")
			os.Exit(1)
		}

		//initialize ldap connection
		err = LdapConnect(config)
		if err != nil {
			os.Exit(1)
		}

		domainControllersQuery(&QueryInfo)
	},
}

// DomainMAQCmd query domain MAQ
var DomainMAQCmd = &cobra.Command{
	Use:   "MAQ",
	Short: "查询域内MAQ",
	Long:  "查询域内MAQ",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := ParseGlobalInfo(cmd)
		if err != nil {
			colors.ErrorPrintln(err)
			colors.ErrorPrintln("Go-ADExec ldap MAQ -h")
			os.Exit(1)
		}

		//initialize ldap connection
		err = LdapConnect(config)
		if err != nil {
			os.Exit(1)
		}

		domainMAQQuery(&QueryInfo)
	},
}

// DomainOUCmd query domain OU
var DomainOUCmd = &cobra.Command{
	Use:   "OU",
	Short: "查询域内OU",
	Long:  "查询域内OU",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := ParseGlobalInfo(cmd)
		if err != nil {
			colors.ErrorPrintln(err)
			colors.ErrorPrintln("Go-ADExec ldap OU -h")
			os.Exit(1)
		}

		//initialize ldap connection
		err = LdapConnect(config)
		if err != nil {
			os.Exit(1)
		}

		domainOUQuery(&QueryInfo)
	},
}

// DomainComputerCmd query domain all Computer
var DomainComputerCmd = &cobra.Command{
	Use:   "DCOMP",
	Short: "查询域内所有机器账户",
	Long:  "查询域内所有机器账户",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := ParseGlobalInfo(cmd)
		if err != nil {
			colors.ErrorPrintln(err)
			colors.ErrorPrintln("Go-ADExec ldap DCOMP -h")
			os.Exit(1)
		}

		//initialize ldap connection
		err = LdapConnect(config)
		if err != nil {
			os.Exit(1)
		}

		domainComputerQuery(&QueryInfo)
	},
}

// DomainUserCmd query domain all user
var DomainUserCmd = &cobra.Command{
	Use:   "DUSER",
	Short: "查询域内所有用户账户",
	Long:  "查询域内所有用户账户",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := ParseGlobalInfo(cmd)
		if err != nil {
			colors.ErrorPrintln(err)
			colors.ErrorPrintln("Go-ADExec ldap DUSER -h")
			os.Exit(1)
		}

		//initialize ldap connection
		err = LdapConnect(config)
		if err != nil {
			os.Exit(1)
		}

		domainUserQuery(&QueryInfo)
	},
}

// DomainGroupCmd query domain all group
var DomainGroupCmd = &cobra.Command{
	Use:   "DGROUP",
	Short: "查询域内所有组",
	Long:  "查询域内所有组",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := ParseGlobalInfo(cmd)
		if err != nil {
			colors.ErrorPrintln(err)
			colors.ErrorPrintln("Go-ADExec ldap DGROUP -h")
			os.Exit(1)
		}

		//initialize ldap connection
		err = LdapConnect(config)
		if err != nil {
			os.Exit(1)
		}

		domainGroupQuery(&QueryInfo)
	},
}

// DomainBaseInfoCmd query domain all group
var DomainBaseInfoCmd = &cobra.Command{
	Use:   "BaseInfo",
	Short: "查询域的基本信息",
	Long:  "查询域的基本信息",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := ParseGlobalInfo(cmd)
		if err != nil {
			colors.ErrorPrintln(err)
			colors.ErrorPrintln("Go-ADExec ldap BaseInfo -h")
			os.Exit(1)
		}

		//initialize ldap connection
		err = LdapConnect(config)
		if err != nil {
			os.Exit(1)
		}

		domainBaseInfoQuery(&QueryInfo)
	},
}
