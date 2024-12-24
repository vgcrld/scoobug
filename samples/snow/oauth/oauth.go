package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/vgcrld/scoobug/samples/snow/oauth/helps"
)

// used to hold the authentication token
var accessTokens *helps.Tokens

// l is a logger
var l *log.Logger

// main function
func main() {

	l = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	c := helps.Creds{}
	accessTokens = getTokens(c.GetCreds())
	// loop 4 times
	for i := 0; i < 4; i++ {
		if accessTokens.Valid {
			createIncident(*accessTokens, c.GetCreds(), helps.SnowReq)
		} else {
			l.Println("Token is invalid, code is: ", accessTokens.RequestCode)
		}
	}

}

func getTokens(c helps.Creds) *helps.Tokens {
	credsString := fmt.Sprintf(
		"grant_type=%s&client_id=%s&client_secret=%s&username=%s&password=%s",
		c.GrantType, c.ClientID, c.ClientSecret, c.Username, c.Password,
	)
	payload := strings.NewReader(credsString)
	req, _ := http.NewRequest("POST", c.Oauth_url, payload)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, _ := http.DefaultClient.Do(req)
	l.Println("Fetch tokens: ", helps.StatusCodes[res.StatusCode])
	body, _ := io.ReadAll(res.Body)
	var t helps.Tokens
	json.Unmarshal(body, &t)
	t.RequestCode = res.StatusCode
	if res.StatusCode != 200 {
		t.Valid = false
	} else {
		t.Valid = true

	}

	return &t
}

func createIncident(a helps.Tokens, c helps.Creds, snowReq map[string]string) {
	parmsJson, _ := json.Marshal(snowReq)
	parmsReader := strings.NewReader(string(parmsJson))
	req, _ := http.NewRequest("POST", c.Incident_url, parmsReader)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+a.AccessToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		l.Println("Error making request:", err)
		return
	}
	defer res.Body.Close()
	// print the body in json
	body, _ := io.ReadAll(res.Body)
	// marshal the body into helps.IncidentResponse
	var incident helps.IncidentResponse
	json.Unmarshal(body, &incident)
	l.Println("Incident: ", helps.StatusCodes[res.StatusCode])
	l.Println("Number: ", incident.Result.Number)
	l.Println("SysID: ", incident.Result.SysID)
}
