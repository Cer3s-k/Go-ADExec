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
	"github.com/spf13/cobra"
	"strings"
)

type LoginInfo struct {
	DomainName string
	UserName   string
	UserPass   string
	UserHash   string
	GssApi     string
	SSLCon     bool
	BaseDN     string
	Output     string
}

// GlobalLoginInfo GlobalLogin global login struct
var GlobalLoginInfo = LoginInfo{}

// initialize the LoginInfo structure and return the content
func parseGlobalInfo(cmd *cobra.Command) (config *LoginInfo, err error) {

	domainName, err := cmd.Flags().GetString("domain")
	if err != nil {
		colors.PrintErrorf("Failed to parse --domainName-- flag %s", err)
		return nil, err
	}
	if domainName == "" {
		return nil, errors.New("domain name is not specified")
	} else {
		GlobalLoginInfo.DomainName = domainName
	}

	userName, err := cmd.Flags().GetString("username")
	if err != nil {
		colors.PrintErrorf("Failed to parse --username-- flag %s", err)
		return nil, err
	}
	if userName == "" {
		return nil, errors.New("username is not specified")
	} else {
		GlobalLoginInfo.UserName = userName
	}

	userPass, err := cmd.Flags().GetString("password")
	if err != nil {
		colors.PrintErrorf("Failed to parse --password-- flag %s", err)
		return nil, err
	}
	if userPass == "" {
		return nil, errors.New("password is not specified")
	} else {
		GlobalLoginInfo.UserPass = userPass
	}

	userHash, err := cmd.Flags().GetString("hashes")
	if err != nil {
		colors.PrintErrorf("Failed to parse --hash-- flag %s", err)
		return nil, err
	}
	if userHash != "" {
		GlobalLoginInfo.UserHash = userHash
	}

	gssApi, err := cmd.Flags().GetString("gssapi")
	if err != nil {
		colors.PrintErrorf("Failed to parse --gssapi-- flag %s", err)
		return nil, err
	}
	if gssApi != "" {
		GlobalLoginInfo.GssApi = gssApi
	}

	sslCon, err := cmd.Flags().GetBool("ssl")
	if err != nil {
		colors.PrintErrorf("Failed to parse --ssl-- flag %s", err)
		return nil, err
	}
	if sslCon != false {
		GlobalLoginInfo.SSLCon = sslCon
	}

	domainNameArr := strings.Split(domainName, ".")
	baseDN, err := cmd.Flags().GetString("basedn")
	if err != nil {
		colors.PrintErrorf("Failed to parse --base dn-- flag %s", err)
		return nil, err
	}
	if baseDN == "" {
		baseDN = fmt.Sprintf("dc=%s", strings.Join(domainNameArr, ",dc="))
		GlobalLoginInfo.BaseDN = baseDN
	} else {
		GlobalLoginInfo.BaseDN = baseDN
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		colors.PrintErrorf("Failed to parse --export-- flag %s", err)
		return nil, err
	}
	if output != "" {
		GlobalLoginInfo.Output = output
	}

	//format domain username
	if !strings.Contains(userName, "@") && !strings.Contains(userName, "\\") {
		userName = fmt.Sprintf("%s@%s", userName, domainName)
		GlobalLoginInfo.UserName = userName
	}

	return &GlobalLoginInfo, nil
}
