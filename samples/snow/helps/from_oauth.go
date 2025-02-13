package helps

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
