package ldapaction

import "Go-ADExec/colors"

// query for Domain Admins
func domainAdminQuery(query *queryConfig) {
	filter := "(&(sAMAccountName=Domain Admins))"
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
