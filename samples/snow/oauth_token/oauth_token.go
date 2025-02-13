package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/vgcrld/scoobug/samples/snow/helps"
	"gopkg.in/yaml.v2"
)

var (
	CLIENT_ID      string
	CLIENT_SECRET  string
	USERNAME       string
	PASSWORD       string
	GRANT_TYPE     string
	SNOW_OAUTH_URL string
	SNOW_BASE_URL  string
	BEARER_TOKEN   string
	REFRESH_TOKEN  string
)

//go:embed .config.yaml
var configTstYaml string

// init function to load config and response
func init() {

	// Load config from YAML
	var config helps.Config
	yaml.Unmarshal([]byte(configTstYaml), &config)
	CLIENT_ID = config.ClientID
	CLIENT_SECRET = config.ClientSecret
	USERNAME = config.Username
	PASSWORD = config.Password
	GRANT_TYPE = config.GrantType
	SNOW_OAUTH_URL = config.SnowOauthUrl
	SNOW_BASE_URL = config.SnowBaseUrl

}

func main() {

	tokens, err := getOAuthToken()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	BEARER_TOKEN = tokens.AccessToken
	REFRESH_TOKEN = tokens.RefreshToken
	writeResponseToFile(tokens)
	getUsersAndGroups()

}

func getUsersAndGroups() {
	url := SNOW_BASE_URL + "/api/now/table/sys_user_grmember?sysparm_display_value=true&sysparm_fields=user%2Cgroup&sysparm_exclude_reference_link=true&sysparm_limit=999999999"
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
	fmt.Println(string(body))
}

// fetch OAuth token and write to file
func getOAuthToken() (*helps.Response, error) {
	log.Default().Println("Fetching OAuth token...")
	encodeLoad := "grant_type=" + url.QueryEscape(GRANT_TYPE) +
		"&client_id=" + url.QueryEscape(CLIENT_ID) +
		"&client_secret=" + url.QueryEscape(CLIENT_SECRET) +
		"&username=" + url.QueryEscape(USERNAME) +
		"&password=" + url.QueryEscape(PASSWORD)
	payload := strings.NewReader(encodeLoad)
	req, _ := http.NewRequest("POST", SNOW_OAUTH_URL, payload)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, _ := http.DefaultClient.Do(req)
	defer func() { res.Body.Close() }()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}
	var tokenResponse helps.Response
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		log.Println("Error unmarshaling response:", err)
		return nil, err
	}
	log.Default().Println("OAuth token fetched successfully")
	return &tokenResponse, nil
}

// write response to file
func writeResponseToFile(tokens *helps.Response) {
	file, err := os.Create(".response.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(tokens)
	if err != nil {
		fmt.Println("Error encoding tokens to JSON:", err)
		return
	}
}
