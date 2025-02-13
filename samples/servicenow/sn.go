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

	"github.com/vgcrld/scoobug/samples/servicenow/sn"
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

func main() {

	tokens, err := getOAuthToken()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	writeResponseToFile(tokens)
	// getUsersAndGroups()

}

// fetch OAuth token and write to file
func getOAuthToken() (*sn.Response, error) {
	log.Default().Println("Fetching OAuth token...")
	encodeLoad := "grant_type=" + url.QueryEscape(sn.Cfg.GrantType) +
		"&client_id=" + url.QueryEscape(sn.Cfg.ClientID) +
		"&client_secret=" + url.QueryEscape(sn.Cfg.ClientSecret) +
		"&username=" + url.QueryEscape(sn.Cfg.Username) +
		"&password=" + url.QueryEscape(sn.Cfg.Password)
	payload := strings.NewReader(encodeLoad)
	req, _ := http.NewRequest("POST", sn.Cfg.SnowOauthUrl, payload)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	defer func() { res.Body.Close() }()
	if err != nil {
		log.Println("Error making request:", err)
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}
	var tokenResponse sn.Response
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		log.Println("Error unmarshaling response:", err)
		return nil, err
	}
	log.Default().Println("OAuth token fetched successfully")
	return &tokenResponse, nil
}

// write response to file
func writeResponseToFile(tokens *sn.Response) {
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
