package main

import (
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"gopkg.in/yaml.v2"
)

var (
	CLIENT_ID     string
	CLIENT_SECRET string
	USERNAME      string
	PASSWORD      string
	GRANT_TYPE    string
)

type Config struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	GrantType    string `yaml:"grant_type"`
}

//go:embed .config.tst.yaml
var configTstYaml string

func init() {

	var config Config
	yaml.Unmarshal([]byte(configTstYaml), &config)
	CLIENT_ID = config.ClientID
	CLIENT_SECRET = config.ClientSecret
	USERNAME = config.Username
	PASSWORD = config.Password
	GRANT_TYPE = config.GrantType
}

func main() {

	snowUrl := "https://dllgroupdevtst.service-now.com/oauth_token.do"

	encodeLoad := "grant_type=" + url.QueryEscape(GRANT_TYPE) + "&client_id=" + url.QueryEscape(CLIENT_ID) + "&client_secret=" + url.QueryEscape(CLIENT_SECRET) + "&username=" + url.QueryEscape(USERNAME) + "&password=" + url.QueryEscape(PASSWORD)
	payload := strings.NewReader(encodeLoad)

	req, _ := http.NewRequest("POST", snowUrl, payload)

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer func() { res.Body.Close() }()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println(string(body))
}
