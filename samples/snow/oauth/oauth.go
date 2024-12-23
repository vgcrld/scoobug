package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// creds struct contains the necessary credentials for authenticating with the
type creds struct {
	oauth_url    string
	incident_url string
	grantType    string
	clientID     string
	clientSecret string
	username     string
	password     string
}

// token struct contains the fields of the authentication token returned by the
type token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

// used to hold the authentication token
var access *token

// access_creds contains the credentials for authenticating
var myCreds creds = creds{
	oauth_url:    "https://dev256710.service-now.com//oauth_token.do",
	incident_url: "https://dev256710.service-now.com/api/now/table/incident",
	grantType:    "password",
	clientID:     "180d976197a94fdfb978816336ed123c",
	clientSecret: "controlm123",
	username:     "admin",
	password:     "KqdT2U2a%5EoR%2B",
}

// l is a logger
var l *log.Logger

// main function
func main() {

	l = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	access := getTokens(myCreds)
	l.Println("Tokens aquired, will expire in", access.ExpiresIn, "seconds")

}

func getTokens(c creds) *token {
	credsString := fmt.Sprintf(
		"grant_type=%s&client_id=%s&client_secret=%s&username=%s&password=%s",
		c.grantType, c.clientID, c.clientSecret, c.username, c.password,
	)
	payload := strings.NewReader(credsString)
	req, _ := http.NewRequest("POST", c.oauth_url, payload)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, _ := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(res.Body)
	var t token
	json.Unmarshal(body, &t)
	return &t
}

func createIncident(t token, c creds) {
	//body of parms in json
	parms := map[string]string{
		"short_description": "This is a test incident",
		"urgency":           "2",
		"impact":            "2",
		"category":          "software",
		"subcategory":       "os",
		"assignment_group":  "software",
		"assigned_to":       "admin",
	}
	parmsJson, _ := json.Marshal(parms)
	fmt.Println(string(parmsJson))
	parmsReader := strings.NewReader(string(parmsJson))
	req, _ := http.NewRequest("POST", c.incident_url, parmsReader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+access.AccessToken)
	res, _ := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}
