package cmd

import (
	"Go-ADExec/colors"
	"crypto/tls"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/go-ldap/ldap/v3/gssapi"
	"strings"
)

func LdapConnect(globalLogin *LdapInfo) (err error) {
	var conn *ldap.Conn
	var sspiConn *gssapi.SSPIClient

	if globalLogin.GssApi != "" {
		sspiConn, err = gssapi.NewSSPIClient()
		if err != nil {
			colors.PrintError(err)
			return err
		}
	}

	//no use TLS for SSL connections
	if !globalLogin.SSLCon {
		colors.PrintSuccessf("Trying to connecting server ldap://%s:389", globalLogin.Domain)
		conn, err = ldap.Dial("tcp", fmt.Sprintf("%s:389", globalLogin.Domain))
		if err != nil {
			colors.PrintError(err)
			return err
		}
	} else {
		// use TLS for SSL connections
		colors.PrintSuccessf("Trying to connecting server ldaps://%s:636", globalLogin.Domain)
		conn, err = ldap.DialTLS("tcp", fmt.Sprintf("%s:636", globalLogin.Domain),
			&tls.Config{InsecureSkipVerify: true})
		if err != nil {
			colors.PrintError(err)
			return err
		}
	}

	if globalLogin.Hash == "" && globalLogin.Pass != "" {
		colors.PrintSuccessf("Trying to binding server with password")
		colors.PrintSuccessf("Domain Name: %s", globalLogin.Domain)
		colors.PrintSuccessf("username:    %s", globalLogin.User)
		colors.PrintSuccessf("password:    %s", globalLogin.Pass)

		err = conn.Bind(globalLogin.User, globalLogin.Pass)
		if err != nil {
			colors.PrintError(err)
			return err
		}
	} else if globalLogin.User != "" && globalLogin.GssApi == "" {
		req := &ldap.NTLMBindRequest{
			Domain:             globalLogin.Domain,
			Username:           strings.Split(globalLogin.User, "@")[0],
			Hash:               globalLogin.Hash,
			AllowEmptyPassword: true,
			Controls:           nil,
		}

		colors.PrintSuccess("Trying to binding server with hash")
		colors.PrintSuccess("username:  %s", strings.Split(globalLogin.User, "@")[0])
		colors.PrintSuccess("ntlm-hash: %s", globalLogin.Hash)

		_, err = conn.NTLMChallengeBind(req)
		if err != nil {
			colors.PrintError(err)
			return err
		}
	} else {
		colors.PrintSuccess("Trying to binging server with current token")

		err = conn.GSSAPIBind(sspiConn, fmt.Sprintf("ldap/%s", globalLogin.GssApi), "")
		if err != nil {
			colors.PrintError(err)
			return err
		}
	}

	colors.PrintSuccess("Binding success")
	globalLogin.Connect = conn

	searchRequest := ldap.NewSearchRequest(globalLogin.BaseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, "(objectClass=user)", []string{"sAMAccountName"}, nil)
	search, err := conn.Search(searchRequest)
	if err != nil {
		colors.PrintError("error: ", err)
		return err
	}
	var sAMAccountName []string = make([]string, 0)
	for j := range search.Entries {
		sAMAccountName = append(sAMAccountName, search.Entries[j].Attributes[0].Values[0])
	}
	for _, sam := range sAMAccountName {
		colors.PrintSuccess(sam)
	}
	return err
}
