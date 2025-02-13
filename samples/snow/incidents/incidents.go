package main

/*
	The program creates an incident in ServiceNow using OAuth authentication to obtain
	an access token and then sending the incident details via an HTTP POST request.

*/
import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/vgcrld/scoobug/samples/snow/helps"
)

// used to hold the authentication token
var accessTokens *helps.Tokens

// l is a logger
var l *log.Logger

// main function
func main() {

	resp, err := getJsonFromFile("../oauth_token/.response.json")
	fmt.Println(resp, err)

	desc := flag.String("desc", "", "Description of the incident")
	flag.Parse()
	if *desc == "" {
		fmt.Println("Description is required")
		os.Exit(1)
	}
	l = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	c := helps.Creds{}
	accessTokens = getTokens(c.GetCreds())
	sr := helps.SnowReq
	sr["short_description"] = *desc
	if accessTokens.Valid {
		createIncident(*accessTokens, c.GetCreds(), sr)
	} else {
		l.Println("Token is invalid, code is: ", accessTokens.RequestCode)
	}

}

func getJsonFromFile(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data map[string]string
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		panic("can't get " + filePath + ":" + err.Error())
	}

	return data, nil
}

func getTokens(c helps.Creds) *helps.Tokens {
	credsString := fmt.Sprintf(
		"grant_type=%s&client_id=%s&client_secret=%s&username=%s&password=%s",
		c.GrantType, c.ClientID, c.ClientSecret, c.Username, c.Password,
	)
	payload := strings.NewReader(credsString)
	req, _ := http.NewRequest("POST", c.OauthUrl, payload)
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
	req, _ := http.NewRequest("POST", c.IncidentUrl, parmsReader)
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
