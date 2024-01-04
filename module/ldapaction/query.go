package ldapaction

import "Go-ADExec/colors"

// query for Domain Admins
func domainAdminQuery(query *queryConfig) {
	filter := "(&(objectCategory=group)(adminCount=1))"
	attributes := []string{"member"}
	query.Attr.Filter = filter
	query.Attr.Attributes = attributes
	res, err := querySearch(query)
	if err != nil {
		colors.ErrorPrintln(err)
		return
	}
	queryPrint(res)
}

// query for domain controllers
func domainControllersQuery(query *queryConfig) {
	filter := "(&(objectCategory=computer)(|(primaryGroupID=521)(primaryGroupID=516)))"
	attributes := []string{"SAMAccountName", "lastLogon"}
	query.Attr.Filter = filter
	query.Attr.Attributes = attributes
	res, err := querySearch(query)
	if err != nil {
		colors.ErrorPrintln(err)
		return
	}
	queryPrint(res)
}

// query for domain MAQ
func domainMAQQuery(query *queryConfig) {
	filter := "(objectClass=domain)"
	attributes := []string{"ms-DS-MachineAccountQuota"}
	query.Attr.Filter = filter
	query.Attr.Attributes = attributes
	res, err := querySearch(query)
	if err != nil {
		colors.ErrorPrintln(err)
		return
	}
	queryPrint(res)
}

// query for domain OU
func domainOUQuery(query *queryConfig) {
	filter := "(&(objectCategory=organizationalUnit)(ou=*))"
	attributes := []string{"distinguishedName"}
	query.Attr.Filter = filter
	query.Attr.Attributes = attributes
	res, err := querySearch(query)
	if err != nil {
		colors.ErrorPrintln(err)
		return
	}
	queryPrint(res)
}

// query for all Computer
func domainComputerQuery(query *queryConfig) {
	filter := "(objectCategory=Computer)"
	attributes := []string{"SAMAccountName", "lastLogon"}
	query.Attr.Filter = filter
	query.Attr.Attributes = attributes
	res, err := querySearch(query)
	if err != nil {
		colors.ErrorPrintln(err)
		return
	}
	queryPrint(res)
}

// query for all user
func domainUserQuery(query *queryConfig) {
	filter := "(objectClass=user)"
	attributes := []string{"SAMAccountName", "lastLogon"}
	query.Attr.Filter = filter
	query.Attr.Attributes = attributes
	res, err := querySearch(query)
	if err != nil {
		colors.ErrorPrintln(err)
		return
	}
	queryPrint(res)
}

// query for all group
func domainGroupQuery(query *queryConfig) {
	filter := "(objectCategory=group)"
	attributes := []string{"distinguishedName"}
	query.Attr.Filter = filter
	query.Attr.Attributes = attributes
	res, err := querySearch(query)
	if err != nil {
		colors.ErrorPrintln(err)
		return
	}
	queryPrint(res)
}

// query for domain baseInfo
func domainBaseInfoQuery(query *queryConfig) {
	attributes := []string{"distinguishedName"}
	query.Attr.Attributes = attributes

	//computer count
	filter1 := "(objectClass=computer)"
	query.Attr.Filter = filter1
	res1, err1 := querySearch(query)

	// user count
	filter2 := "(objectClass=user)"
	query.Attr.Filter = filter2
	res2, err2 := querySearch(query)

	if err1 != nil || err2 != nil {
		colors.ErrorPrintln(err1)
		colors.ErrorPrintln(err2)
		return
	}
	colors.SuccessPrintf("域内机器数量: %d\n", len(res1))
	colors.SuccessPrintf("域内用户数量: %d\n", len(res2))
	//queryPrint(res)
}
