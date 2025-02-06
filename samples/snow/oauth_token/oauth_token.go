package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

var (
	CLIENT_ID     string
	CLIENT_SECRET string
	USERNAME      string
	PASSWORD      string
	GRANT_TYPE    string
	SNOWURL       string

	ACCESS_TOKEN  string
	REFRESH_TOKEN string
)

type Config struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	GrantType    string `yaml:"grant_type"`
	SnowUrl      string `yaml:"snow_url"`
}

type Response struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

//go:embed .config.yaml
var configTstYaml string

//go:embed .response.json
var responseJson string

func init() {

	// Load config from YAML
	var config Config
	yaml.Unmarshal([]byte(configTstYaml), &config)
	CLIENT_ID = config.ClientID
	CLIENT_SECRET = config.ClientSecret
	USERNAME = config.Username
	PASSWORD = config.Password
	GRANT_TYPE = config.GrantType
	SNOWURL = config.SnowUrl

	// Load response from JSON
	var response Response
	json.Unmarshal([]byte(responseJson), &response)
	ACCESS_TOKEN = response.AccessToken
	REFRESH_TOKEN = response.RefreshToken

}

func main() {

	encodeLoad := "grant_type=" + url.QueryEscape(GRANT_TYPE) + "&client_id=" + url.QueryEscape(CLIENT_ID) + "&client_secret=" + url.QueryEscape(CLIENT_SECRET) + "&username=" + url.QueryEscape(USERNAME) + "&password=" + url.QueryEscape(PASSWORD)
	payload := strings.NewReader(encodeLoad)

	req, _ := http.NewRequest("POST", SNOWURL, payload)

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer func() { res.Body.Close() }()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	writeResponseToFile(body)

}

func writeResponseToFile(body []byte) {
	filePath := ".response.json"
	err := os.WriteFile(filePath, body, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Response written to", filePath)
}
