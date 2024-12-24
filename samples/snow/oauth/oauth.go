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
type tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Valid        bool
	RequestCode  int
}

// used to hold the authentication token
var accessTokens *tokens

// access_creds contains the credentials for authenticating
var myCreds creds = creds{
	oauth_url:    "https://dev256710.service-now.com//oauth_token.do",
	incident_url: "https://dev256710.service-now.com/api/now/table/incident",
	grantType:    "password",
	clientID:     "2eda594ac70345fc80a47c1e5b4c2424",
	clientSecret: "controlm123",
	username:     "admin",
	password:     "KqdT2U2a%5EoR%2B",
	// password is encoded and needs to be. (same issue nithila saw)
	// the actual password is 'KqdT2U2a^oR+'  ^ and + are special characters and need to be encoded
}

// l is a logger
var l *log.Logger

// when creating this you have to be careful with the fields you are sending
// it will trigger a business rule if you are not careful and fail with a 403
// this failed on a 403 when trying to assign it to a user "admin"
var snowReq = map[string]string{
	"short_description": "This is a test incident",
	"urgency":           "2",
	"impact":            "2",
	"category":          "software",
	"subcategory":       "os",
	"assignment_group":  "software",
	// "assigned_to":       "admin",
}

// main function
func main() {

	l = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	accessTokens = getTokens(myCreds)
	if accessTokens.Valid {
		createIncident(*accessTokens, myCreds, snowReq)
	} else {
		l.Println("Token is invalid, code is: ", accessTokens.RequestCode)
	}

}

func getTokens(c creds) *tokens {
	credsString := fmt.Sprintf(
		"grant_type=%s&client_id=%s&client_secret=%s&username=%s&password=%s",
		c.grantType, c.clientID, c.clientSecret, c.username, c.password,
	)
	payload := strings.NewReader(credsString)
	req, _ := http.NewRequest("POST", c.oauth_url, payload)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, _ := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(res.Body)
	var t tokens
	json.Unmarshal(body, &t)
	t.RequestCode = res.StatusCode
	if res.StatusCode != 200 {
		t.Valid = false
	} else {
		t.Valid = true

	}

	return &t
}

func createIncident(a tokens, c creds, snowReq map[string]string) {
	l.Println("Creating incident")
	parmsJson, _ := json.Marshal(snowReq)
	parmsReader := strings.NewReader(string(parmsJson))
	req, _ := http.NewRequest("POST", c.incident_url, parmsReader)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+a.AccessToken)

	// Log the request
	l.Println("Request URL:", req.URL)
	l.Println("Request Headers:", req.Header)
	l.Println("Request Body:", string(parmsJson))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		l.Println("Error making request:", err)
		return
	}
	defer res.Body.Close()

	// Log the response
	body, _ := io.ReadAll(res.Body)
	// l.Println("Response Status:", res.Status)
	// l.Println("Response Headers:", res.Header)
	l.Println("Response Body:", string(body))

	fmt.Println("Status code:", res.StatusCode)
}
