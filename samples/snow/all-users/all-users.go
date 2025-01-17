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
	Result []UserGroup `json:"result"`
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
	getGroups(&response, &groups)

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

func getGroups(response *Response, groups *map[string]int) {
	url := "https://dllgroupdevtst.service-now.com/api/now/table/sys_user_grmember?sysparm_display_value=true&sysparm_fields=user%2Cgroup&sysparm_exclude_reference_link=true&sysparm_limit=999999999"
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
	err := json.Unmarshal(body, &response)
	totalFound := 0
	if err == nil && resCode == 200 {
		totalFound = len(response.Result)
	}
	log.Printf("Getting groups StatusCode=%d, records=%d)", resCode, totalFound)
	addToResponse(response, groups)
}

func addToResponse(response *Response, groups *map[string]int) {
	for _, result := range response.Result {
		g := result.Group
		if (*groups)[g] == 0 {
			(*groups)[g] = 0
		}
		(*groups)[g]++
	}

}
