// Package cmd
/*
This file is responsible for defining the structure of ldap connection information
and initializing it based on the parameters passed in.
*/
package cmd

import (
	"Go-ADExec/colors"
	"errors"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/spf13/cobra"
	"strings"
)

type LdapInfo struct {
	LdapServer string
	LdapIP     string
	Domain     string
	User       string
	Pass       string
	Hash       string
	Connect    *ldap.Conn
	GssApi     string
	SSLCon     bool
	BaseDN     string
	Output     string
}

// GlobalLoginInfo GlobalLogin global login struct
var GlobalLoginInfo = LdapInfo{}

// initialize the LoginInfo structure and return the content
func parseGlobalInfo(cmd *cobra.Command) (config *LdapInfo, err error) {

	domainName, err := cmd.Flags().GetString("domain")
	if err != nil {
		colors.PrintErrorf("Failed to parse --domain flag %s", err)
		return nil, err
	}
	if domainName == "" {
		return nil, errors.New("domain name is not specified")
	} else {
		GlobalLoginInfo.Domain = strings.Trim(domainName, "'")
	}

	userName, err := cmd.Flags().GetString("username")
	if err != nil {
		colors.PrintErrorf("Failed to parse --username flag %s", err)
		return nil, err
	}
	if userName == "" {
		return nil, errors.New("username is not specified")
	} else {
		GlobalLoginInfo.User = strings.Trim(userName, "'")
	}

	userPass, err := cmd.Flags().GetString("password")
	if err != nil {
		colors.PrintErrorf("Failed to parse --password flag %s", err)
		return nil, err
	}
	if userPass == "" {
		return nil, errors.New("password is not specified")
	} else {
		GlobalLoginInfo.Pass = strings.Trim(userPass, "'")
	}

	userHash, err := cmd.Flags().GetString("hashes")
	if err != nil {
		colors.PrintErrorf("Failed to parse --hashes flag %s", err)
		return nil, err
	}
	if userHash != "" {
		GlobalLoginInfo.Hash = strings.Trim(userHash, "'")
	}

	gssApi, err := cmd.Flags().GetString("gssapi")
	if err != nil {
		colors.PrintErrorf("Failed to parse --gssapi flag %s", err)
		return nil, err
	}
	if gssApi != "" {
		GlobalLoginInfo.GssApi = strings.Trim(gssApi, "'")
	}

	sslCon, err := cmd.Flags().GetBool("ssl")
	if err != nil {
		colors.PrintErrorf("Failed to parse --ssl flag %s", err)
		return nil, err
	}
	if sslCon != false {
		GlobalLoginInfo.SSLCon = sslCon
	}

	domainNameArr := strings.Split(domainName, ".")
	baseDN, err := cmd.Flags().GetString("basedn")
	if err != nil {
		colors.PrintErrorf("Failed to parse --basedn flag %s", err)
		return nil, err
	}
	if baseDN == "" {
		baseDN = fmt.Sprintf("dc=%s", strings.Join(domainNameArr, ",dc="))
		GlobalLoginInfo.BaseDN = baseDN
	} else {
		GlobalLoginInfo.BaseDN = strings.Trim(baseDN, "'")
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		colors.PrintErrorf("Failed to parse --output flag %s", err)
		return nil, err
	}
	if output != "" {
		GlobalLoginInfo.Output = strings.Trim(output, "'")
	}

	//format domain username
	if !strings.Contains(userName, "@") && !strings.Contains(userName, "\\") {
		userName = fmt.Sprintf("%s@%s", strings.Trim(userName, "'"), domainName)
		GlobalLoginInfo.User = userName
	}

	return &GlobalLoginInfo, nil
}
