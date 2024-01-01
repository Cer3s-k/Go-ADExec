package ldapaction

import (
	"Go-ADExec/colors"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"strings"
)

// CustomSearch custom ldap search
func CustomSearch() {

	QueryInfo.Attr.Attributes = append(QueryInfo.Attr.Attributes, "distinguishedName")
	QueryInfo.Attr.Attributes = removeDuplicate(QueryInfo.Attr.Attributes)

	res, err := querySearch(&QueryInfo)
	if err != nil {
		colors.ErrorPrintln(err)
		return
	}

	queryPrint(res)

}

// global ldap query function
func querySearch(queryInfo *queryConfig) ([]*ldap.Entry, error) {
	colors.InfoPrintf("Search info:\n"+
		"    base dn:   %s\n"+
		"    filter:    %s\n"+
		"    attribute: %s\n", queryInfo.Global.BaseDN, queryInfo.Attr.Filter, queryInfo.Attr.Attributes)

	searchRequest := ldap.NewSearchRequest(
		queryInfo.Global.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		queryInfo.Attr.Filter,
		queryInfo.Attr.Attributes,
		nil,
	)

	// perform a paged search
	sr, err := queryInfo.Global.Connect.SearchWithPaging(searchRequest, 1000)
	if err != nil {
		return nil, err
	}

	return sr.Entries, nil
}

// global ldap query output
func queryPrint(queryResult []*ldap.Entry) {

	var successResult string
	var normalResult string
	var attResult strings.Builder
	var saveResultStr = "result.txt"
	var isWrite = false

	// judge output
	for _, entry := range queryResult {
		successResult = fmt.Sprintf("%s\n", entry.DN)

		for _, v := range entry.Attributes {
			switch v.Name {
			case "nTSecurityDescriptor", "msDS-AllowedToActOnBehalfOfOtherIdentity":
				colors.SuccessPrintln("this is s acl related")
				//sr, err := sddl.NewSecurityDescriptor(v.ByteValues[0])
				//if err != nil {
				//	log.PrintErrorf("%s\n%s\n", "resolve nTSecurityDescriptor error:", err.Error())
				//	return
				//}
				//log.PrintDebugf("dump nTSecurityDescriptor string: \n%s\n", sr.DataToString(v.ByteValues[0]))
				//
				//var endResult strings.Builder
				//
				//if sr.OwnerSid.Value != nil {
				//	endResult.WriteString(fmt.Sprintf("[OwnerSid: %s]", sr.OwnerSid.Value.(string)))
				//}
				//
				//if sr.GroupSid.Value != nil {
				//	endResult.WriteString(fmt.Sprintf("[GroupSid: %s]", sr.GroupSid.Value.(string)))
				//}
				//
				//if sr.Dacl.AclSize.Value != 0 {
				//	for _, ace := range sr.Dacl.Aces {
				//		aceMaskString, err := ace.AceMask.GetAceMaskString()
				//		if err != nil {
				//			log.PrintErrorf("get ace mask string error: %s", err)
				//			os.Exit(-2)
				//		}
				//
				//		endResult.WriteString(fmt.Sprintf("[[Ace Mask: %s]", aceMaskString))
				//
				//		if ace.ObjectType != nil {
				//			endResult.WriteString(fmt.Sprintf("[ObjectType: %s]", ace.ObjectType.Value.(string)))
				//		}
				//
				//		if ace.InheritedObjectType != nil {
				//			endResult.WriteString(fmt.Sprintf("[InheritedObjectType: %s]", ace.InheritedObjectType.Value.(string)))
				//		}
				//
				//		if ace.SID.Value != nil {
				//			endResult.WriteString(fmt.Sprintf("[SID: %s]", ace.SID.Value.(string)))
				//		}
				//
				//		endResult.WriteString("]")
				//	}
				//}
				//
				//result.WriteString(fmt.Sprintf("    %s: %s\n", v.Name, endResult.String()))

			case "lastLogon":
				colors.SuccessPrintln("this is a time related")
				//dateString, err := transform.TimeToString(v.Values[0])
				//if err != nil {
				//	log.PrintErrorf("%s\n%s\n", "resolve lastlogon error: ", err.Error())
				//	return
				//}
				//result.WriteString(fmt.Sprintf("    %s: %s\n", v.Name, dateString))
			case "objectGUID":
				colors.SuccessPrintln("this is GUID related")
				//toString, err := guid.GuidToString(v.ByteValues[0])
				//if err != nil {
				//	log.PrintErrorf("%s\n%s\n", "resolve objectGUID error: ", err.Error())
				//	return
				//}
				//
				//result.WriteString(fmt.Sprintf("    %s: %s\n", v.Name, toString))

			default:
				t := fmt.Sprintf("\n    %s: ", v.Name)
				attResult.WriteString(fmt.Sprintf("    %s: %s\n", v.Name, strings.Join(v.Values, t)))
			}
			//if v.Name == "nTSecurityDescriptor" {
			//	sr, err := sddl.NewSecurityDescriptor(v.ByteValues[0])
			//	if err != nil {
			//		log.PrintErrorf("%s\n%s\n", "resolve nTSecurityDescriptor error:", err.Error())
			//		return
			//	}
			//	resultStrings := sr.DataToString(v.ByteValues[0])
			//	log.PrintDebugf("dump nTSecurityDescriptor string: \n%s\n", resultStrings.String())
			//} else {
			//	result.WriteString(fmt.Sprintf("    %s: %s\n", v.Name, strings.Join(v.Values, " ")))
			//}
		}

		normalResult = attResult.String()
		attResult = strings.Builder{}

		// deal with the results
		if len(queryResult) < 50 && QueryInfo.Global.Output == "" {
			colors.SuccessPrintf(successResult)
			colors.NormalPrintf(normalResult)
		} else {
			isWrite = true
			//log.SaveResultStr = "result.txt"
			//log.PrintWarningf("The number of returned results is too large, the output will be saved in %s", log.SaveResultStr)
			//
			//err := log.SaveResult([]byte(result.String()))
			//if err != nil {
			//	log.PrintErrorf("Save Result error: %s", err.Error())
			//	os.Exit(-2)
			//}
			//
			//log.SaveResultStr = ""
		}
	}

	if isWrite {
		if len(queryResult) > 50 {
			colors.InfoPrintf("too many results are returned. save the results in the %s file instead\n", saveResultStr)
		} else if QueryInfo.Global.Output != "" {
			saveResultStr = QueryInfo.Global.Output
			colors.InfoPrintf("save the results in the %s file\n", saveResultStr)
		}

	}

	colors.InfoPrintf("result count: %d\n", len(queryResult))

	//if log.SaveResultStr != "" {
	//	log.PrintInfof("saving result to %s", log.SaveResultStr)
	//	err := log.SaveResult([]byte(result.String()))
	//	if err != nil {
	//		log.PrintErrorf("Save Result error: %s", err.Error())
	//		os.Exit(-2)
	//	}
	//}

	//for _, entry := range queryResult {
	//	colors.SuccessPrintln(entry.DN)
	//	for _, attribute := range entry.Attributes {
	//		colors.NormalPrintf("    %s: ", attribute.Name)
	//		colors.NormalPrintf("%s\n", attribute.Values[0])
	//	}
	//
	//}

}

// RemoveDuplicate filter duplicate elements by unique characteristics of map primary key
func removeDuplicate(arr []string) []string {
	resArr := make([]string, 0)
	tmpMap := make(map[string]interface{})
	for _, val := range arr {
		//判断主键为val的map是否存在
		if _, ok := tmpMap[val]; !ok {
			resArr = append(resArr, val)
			tmpMap[val] = nil
		}
	}

	return resArr
}