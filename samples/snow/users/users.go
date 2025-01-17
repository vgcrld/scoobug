package main

/*
	this sample code is used to get all users in a group from ServiceNow
	You will need a valid bearer token to access the ServiceNow API

*/

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type UserGroup struct {
	User  string `json:"user"`
	Group string `json:"group"`
}

type Response struct {
	Result    []UserGroup `json:"result"`
	GroupName string
}

var allGroups []string = []string{
	"Americas_Canada",
	"Americas_Des_Moines_Canada_DWH",
	"Americas_LATAM_Delivery",
	"Americas_LATAM_MAC_Support",
	"Americas_NA_OracleCloudSaaS",
	"Americas_Wayne_BO_CCRS",
	"Americas_Wayne_BO",
	"Americas_Wayne_FO",
	"BUS_Qlik_ US",
	"DIGITAL_API",
	"DIGITAL_OFNG",
	"EBS_Business_Process_Management",
	"EBS_Enterprise_Content_Management",
	"EBS_Risk_and_Compliance",
	"EBS_TREASURY",
	"EMEA_VF_Digital_Portals",
	"EMEA_VF_FO",
	"EMEA_VF_Operations",
	"EMEA_VF_OpEx_CreditHub",
	"EMEA_VF_OpEx_Delta",
	"EMEA_VF_OpEx_DocMgmt",
	"EMEA_VF_OpEx_eCS",
	"EMEA_VF_OpEx_Finance",
	"ETS_Document_Management",
	"GOV_CHANGE_CONFIG",
	"GOV_LICENSE_MANAGEMENT",
	"ETS_BACKUP_RESTORE",
	"ETS_DBA_ORACLE",
	"ETS_DBA_SQL",
	"ETS_EVENT_MANAGEMENT",
	"ETS_LINUX",
	"ETS_UCC",
	"ETS_WINDOWS",
	"ETS_WORKLOAD_AUTOMATION",
	"ETS_X86_STORAGE",
	"IT_ANZ",
	"ITSD_APAC",
	"LIM_Italy",
	"LIM_ITALY",
	"NA_DIGITAL_API",
	"NA_DIGITAL_CX",
	"Rock_TAM",
	"ROCK_TAM",
	"SUP_ACN_BI_WN_Support",
	"SUP_ACN_EP_GL_Support",
	"SUP_ACN_GS_CC_Support",
	"SUP_ACN_GS_GL_SUPPORT",
	"SUP_ACN_VF_DM_Support",
	"SUP_ACN_VF_EU_SUPPORT_ACC",
	"SUP_ACN_VF_EU_Support",
	"SUP_ACN_VF_WN_BO_Support",
	"SUP_ACN_VF_WN_BO_SUPPORT",
	"SUP_ACN_VF_WN_FO_Support",
	"SUP_ICC",
	"SUP_TCS_Trapeeze",
}

var BEARER_TOKEN string

func init() {
	if BEARER_TOKEN = os.Getenv("BEARER_TOKEN"); BEARER_TOKEN == "" {
		log.Fatal("BEARER_TOKEN environment variable not set")
	}
}

func main() {

	var response Response
	var groups = make(map[string]int)
	for _, g := range allGroups {
		getGroups(g, &response, &groups)
	}

	file, err := os.Create("groups.csv")
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}
	defer file.Close()

	writer := io.Writer(file)
	writer.Write([]byte("Group,UserCount\n"))
	for group, count := range groups {
		writer.Write([]byte(fmt.Sprintf("%s,%d\n", group, count)))
	}
	log.Println("CSV file created successfully")

}

func getGroups(group string, response *Response, groups *map[string]int) {
	url := "https://dllgroupdevtst.service-now.com/api/now/table/sys_user_grmember?sysparm_display_value=true&sysparm_fields=user%2Cgroup&sysparm_exclude_reference_link=true&sysparm_query=group.name=" + group
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Authorization", "Bearer "+BEARER_TOKEN)
	res, _ := http.DefaultClient.Do(req)
	defer func() { res.Body.Close() }()
	body, _ := io.ReadAll(res.Body)
	resCode := res.StatusCode
	if resCode == 401 {
		log.Println("Unauthorized access. Please check your bearer token.")
		os.Exit(1)
	}
	response.GroupName = group
	err := json.Unmarshal(body, &response)
	totalFound := 0
	if err == nil && resCode == 200 {
		totalFound = len(response.Result)
	}
	log.Printf("Getting groups for: %s (StatusCode: %d) (records: %d)", group, resCode, totalFound)
	addToResponse(response, groups)
}

func addToResponse(response *Response, groups *map[string]int) {
	g := response.GroupName
	(*groups)[g] = len(response.Result)
}
