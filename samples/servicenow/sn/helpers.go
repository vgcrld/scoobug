package sn

import (
	_ "embed"

	"gopkg.in/yaml.v2"
)

//go:embed .config.yaml
var configYaml string

// the main config struct
var Cfg Config

// init function to load config and response
func init() {

	// Load config from YAML
	yaml.Unmarshal([]byte(configYaml), &Cfg)

}

// map[int]string of the status cods and thier meeting for the service now api
var HttpStatusCodes = map[int]string{
	200: "OK",
	201: "Created",
	204: "No Content",
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	415: "Unsupported Media Type",
	500: "Internal Server Error",
}

// creds struct contains the necessary credentials for authenticating with the
type Creds struct {
	OauthUrl     string
	IncidentUrl  string
	GrantType    string
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
}

func (c Creds) GetCreds() Creds {
	return Creds{
		OauthUrl:     "https://dev256710.service-now.com//oauth_token.do",
		IncidentUrl:  "https://dev256710.service-now.com/api/now/table/incident",
		GrantType:    "password",
		ClientID:     "2eda594ac70345fc80a47c1e5b4c2424",
		ClientSecret: "controlm123",
		Username:     "sa-ctm",
		Password:     "0X6]flZ]J07BN={pTN2{>u_kP",
	}
}

// token struct contains the fields of the authentication token returned by the
type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Valid        bool
	RequestCode  int
}

// when creating this you have to be careful with the fields you are sending
// it will trigger a business rule if you are not careful and fail with a 403
// this failed on a 403 when trying to assign it to a user "admin"
var SnowReq = map[string]string{
	"short_description": "This is a test incident",
	"urgency":           "2",
	"impact":            "2",
	"category":          "software",
	"subcategory":       "os",
	"assignment_group":  "software",
	// "assigned_to":       "admin",
}

// IncidentResponse struct contains the fields of the response returned by the ServiceNow API
type IncidentResponse struct {
	Result struct {
		Number           string `json:"number"`
		SysID            string `json:"sys_id"`
		ShortDescription string `json:"short_description"`
		Urgency          string `json:"urgency"`
		Impact           string `json:"impact"`
		Category         string `json:"category"`
		Subcategory      string `json:"subcategory"`
		AssignmentGroup  string `json:"assignment_group"`
		AssignedTo       string `json:"assigned_to"`
	} `json:"result"`
}

type Config struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	GrantType    string `yaml:"grant_type"`
	SnowOauthUrl string `yaml:"snow_oauth_url"`
	SnowBaseUrl  string `yaml:"snow_base_url"`
}

type Response struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserGroup struct {
	User  string `json:"user"`
	Group string `json:"group"`
}

type GroupResponse struct {
	Result []UserGroup `json:"result"`
}
