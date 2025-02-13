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

	"github.com/vgcrld/scoobug/samples/snow/helps"
)

var BEARER_TOKEN string

func init() {
	if BEARER_TOKEN = os.Getenv("BEARER_TOKEN"); BEARER_TOKEN == "" {
		log.Fatal("BEARER_TOKEN environment variable not set")
	}
}

func main() {

	var response helps.GroupResponse
	var groups = make(map[string][]string)
	getGroups(&response, &groups)

	file, err := os.Create("groups.csv")
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}
	defer file.Close()

	writer := io.Writer(file)
	writer.Write([]byte("Group,UserCount\n"))
	for group, names := range groups {
		for _, name := range names {
			writer.Write([]byte(fmt.Sprintf("%s,%s\n", group, name)))
		}
	}
	log.Println("CSV file created successfully")

}

func getGroups(response *helps.GroupResponse, groups *map[string][]string) {
	url := "https://dev256710.service-now.com/api/now/table/sys_user_grmember?sysparm_display_value=true&sysparm_fields=user%2Cgroup&sysparm_exclude_reference_link=true&sysparm_limit=999999999"
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

func addToResponse(response *helps.GroupResponse, groups *map[string][]string) {
	for _, result := range response.Result {
		g := result.Group
		u := result.User
		(*groups)[g] = append((*groups)[g], u)
	}

}
