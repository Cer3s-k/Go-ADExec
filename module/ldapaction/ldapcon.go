//go:build !windows

package ldapaction

import (
	"Go-ADExec/colors"
	"crypto/tls"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"strings"
)

func LdapConnect(globalLogin *LdapInfo) (err error) {

	var conn *ldap.Conn

	//no use TLS for SSL connections
	if !globalLogin.SSLCon {
		colors.InfoPrintf("Trying to connecting server ldap://%s:389\n", globalLogin.Domain)
		conn, err = ldap.Dial("tcp", fmt.Sprintf("%s:389", globalLogin.Domain))
		if err != nil {
			colors.ErrorPrintln(err)
			return err
		}
	} else {
		// use TLS for SSL connections
		colors.InfoPrintf("Trying to connecting server ldaps://%s:636\n", globalLogin.Domain)
		conn, err = ldap.DialTLS("tcp", fmt.Sprintf("%s:636", globalLogin.Domain),
			&tls.Config{InsecureSkipVerify: true})
		if err != nil {
			colors.ErrorPrintln(err)
			return err
		}
	}

	if globalLogin.Pass != "" {
		colors.InfoPrintln("Trying to binding server with password")
		colors.InfoPrintf("Domain Name: %s\n", globalLogin.Domain)
		colors.InfoPrintf("username: %s\n", globalLogin.User)
		colors.InfoPrintf("password: %s\n", globalLogin.Pass)

		err = conn.Bind(globalLogin.User, globalLogin.Pass)
		if err != nil {
			colors.ErrorPrintln(err)
			return err
		}
	} else if globalLogin.Hash != "" {
		req := &ldap.NTLMBindRequest{
			Domain:             globalLogin.Domain,
			Username:           strings.Split(globalLogin.User, "@")[0],
			Hash:               globalLogin.Hash,
			AllowEmptyPassword: true,
			Controls:           nil,
		}

		colors.InfoPrintln("Trying to binding server with hash")
		colors.InfoPrintf("username:  %s\n", strings.Split(globalLogin.User, "@")[0])
		colors.InfoPrintf("ntlm-hash: %s\n", globalLogin.Hash)

		_, err = conn.NTLMChallengeBind(req)
		if err != nil {
			colors.ErrorPrintln(err)
			return err
		}
	}

	colors.InfoPrintln("Binding success")
	globalLogin.Connect = conn

	//searchRequest := ldap.NewSearchRequest(globalLogin.BaseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, "(objectClass=user)", []string{"sAMAccountName"}, nil)
	//search, err := conn.Search(searchRequest)
	//if err != nil {
	//	colors.ErrorPrintln("error: ", err)
	//	return err
	//}
	//var sAMAccountName []string = make([]string, 0)
	//for j := range search.Entries {
	//	sAMAccountName = append(sAMAccountName, search.Entries[j].Attributes[0].Values[0])
	//}
	//for _, sam := range sAMAccountName {
	//	colors.SuccessPrintln(sam)
	//}
	return err
}
