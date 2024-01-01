//go:build windows

package ldapaction

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
			colors.ErrorPrintln(err)
			return err
		}
	}

	//no use TLS for SSL connections
	if !globalLogin.SSLCon {
		colors.InfoPrintf("Trying to connecting server ldapaction://%s:389\n", globalLogin.Domain)
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
	} else {
		colors.InfoPrintln("Trying to binging server with current token")

		err = conn.GSSAPIBind(sspiConn, fmt.Sprintf("ldapaction/%s", globalLogin.GssApi), "")
		if err != nil {
			colors.ErrorPrintln(err)
			return err
		}
	}

	colors.InfoPrintln("Binding success")
	globalLogin.Connect = conn

	return err
}
