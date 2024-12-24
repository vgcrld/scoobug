# OAuth Authentication and Incident Creation in Go

This example demonstrates how to authenticate with an OAuth 2.0 provider and create an incident using the ServiceNow API in Go.

## Code

```go
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

// creds struct contains the necessary credentials for authenticating with the OAuth provider.
type creds struct {
	oauth_url    string
	incident_url string
	grantType    string
	clientID     string
	clientSecret string
	username     string
	password     string
}

// token struct contains the fields of the authentication token returned by the OAuth provider.
type tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

// used to hold the authentication token
var accessTokens *tokens

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
	accessTokens = getTokens(myCreds)
	l.Println("Tokens acquired, will expire in", accessTokens.ExpiresIn, "seconds: ", accessTokens.AccessToken)
	createIncident(*accessTokens, myCreds)
}

// getTokens authenticates with the OAuth provider and returns an authentication token.
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
	return &t
}

// createIncident creates an incident in ServiceNow using the provided authentication token and credentials.
func createIncident(a tokens, c creds) {
	// body of parms in json
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
	req.Header.Add("Authorization", "Bearer "+a.AccessToken)
	res, _ := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}
```

## Explanation

### creds Struct

The `creds` struct contains the necessary credentials for authenticating with the OAuth provider. It includes fields for the OAuth URL, incident URL, grant type, client ID, client secret, username, and password.

### tokens Struct

The `tokens` struct contains the fields of the authentication token returned by the OAuth provider. It includes fields for the access token, refresh token, scope, token type, and expiration time.

### main Function

The `main` function is the entry point of the program. It initializes a logger, calls the `getTokens` function to authenticate with the OAuth provider, logs the acquired tokens, and calls the `createIncident` function to create an incident in ServiceNow.

### getTokens Function

The `getTokens` function authenticates with the OAuth provider and returns an authentication token. It constructs a credentials string, creates an HTTP POST request, and sends the request to the OAuth URL. The response body is read and unmarshaled into a `tokens` struct, which is then returned.

### createIncident Function

The `createIncident` function creates an incident in ServiceNow using the provided authentication token and credentials. It constructs a JSON payload with the incident parameters, creates an HTTP POST request, and sends the request to the incident URL. The response body is read and printed.
