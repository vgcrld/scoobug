package tz

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

// list of all the zones fetched
var ZonesQueried []TimezoneInfo

// list of all the timezones to fetch
var ZonesToFetch = map[string]string{
	"America/Chicago":  "CST",
	"America/New_York": "EST",
	"America/Santiago": "CHL",
	// "Etc/Greenwich":    "GMT",
	"Europe/Paris":     "CET",
	"Europe/Bucharest": "EET",
	"Australia/Sydney": "SYD",
}

const (
	APIKEY string = "b07b8d5d64634589a18909071f2da4e3"
)

type DstInfo struct {
	UtcTime        string `json:"utc_time"`
	Duration       string `json:"duration"`
	Gap            bool   `json:"gap"`
	DateTimeAfter  string `json:"dateTimeAfter"`
	DateTimeBefore string `json:"dateTimeBefore"`
	Overlap        bool   `json:"overlap"`
}

type TimezoneInfo struct {
	Timezone              string  `json:"timezone"`
	TimezoneOffset        int     `json:"timezone_offset"`
	TimezoneOffsetWithDst int     `json:"timezone_offset_with_dst"`
	Date                  string  `json:"date"`
	DateTime              string  `json:"date_time"`
	DateTimeTxt           string  `json:"date_time_txt"`
	DateTimeWti           string  `json:"date_time_wti"`
	DateTimeYmd           string  `json:"date_time_ymd"`
	DateTimeUnix          float64 `json:"date_time_unix"`
	Time24                string  `json:"time_24"`
	Time12                string  `json:"time_12"`
	Week                  int     `json:"week"`
	Month                 int     `json:"month"`
	Year                  int     `json:"year"`
	YearAbbr              string  `json:"year_abbr"`
	IsDst                 bool    `json:"is_dst"`
	DstSavings            int     `json:"dst_savings"`
	DstExists             bool    `json:"dst_exists"`
	DstStart              DstInfo `json:"dst_start"`
	DstEnd                DstInfo `json:"dst_end"`
}

type Error struct {
	Zone string
	Err  error
}

// a list of all errors received
var Errors []Error

// returns the short form of the long name
func (t TimezoneInfo) GetShortName() string {
	return ZonesToFetch[t.Timezone]
}

func (t TimezoneInfo) PrintCtmLine() {
	fmt.Println(t.GetShortName())
}

// print the JSON
func PrintJSON() {
	jsonData, _ := json.MarshalIndent(ZonesQueried, "", "  ")
	fmt.Println(string(jsonData))
}

func QueryAll() error {
	for long := range ZonesToFetch {
		log.Printf("Querying timezone: %s\n", long)
		err := QueryZoneAPI(long)
		if err != nil {
			return err
		}
	}
	return nil
}

// fetch the timezone info from the API
func QueryZoneAPI(longname string) (e error) {

	var ret TimezoneInfo

	url := fmt.Sprintf("https://api.ipgeolocation.io/timezone?apiKey=%s&tz=%s", APIKEY, longname)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println(err)
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	} else if res.StatusCode != 200 {
		return errors.New("the server returned a non-200 status code: " + res.Status)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		log.Println(err)
		Errors = append(Errors, Error{Zone: longname, Err: err})
	}
	ZonesQueried = append(ZonesQueried, ret)

	return err
}
