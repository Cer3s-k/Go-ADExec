package cmd

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

type Connector struct {
	Conn   *ldap.Conn
	Config *LoginInfo
}

func LdapConnect(globalLogin *LoginInfo) (*Connector, error) {
	conn, err := ldap.Dial("tcp", fmt.Sprintf("%s:389", globalLogin.DomainName))
	if err != nil {
		return nil, err
	}

	//ldap connection binding
	err = conn.Bind(globalLogin.UserName, globalLogin.UserPass)
	if err != nil {
		return nil, err
	}
	searchRequest := ldap.NewSearchRequest(globalLogin.BaseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, "objectClass=user", []string{"sAMAccountName"}, nil)
	search, err := conn.Search(searchRequest)
	if err != nil {
		return nil, err
	}
	var sAMAccountName []string = make([]string, 0)
	for j := range search.Entries {
		sAMAccountName = append(sAMAccountName, search.Entries[j].Attributes[0].Values[0])
	}
	fmt.Println("sAMAccountName: ", sAMAccountName)
	return &Connector{Conn: conn, Config: globalLogin}, nil
}
