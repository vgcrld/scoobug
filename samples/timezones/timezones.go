package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

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

func main() {
	tzinfo, err := getZone("America/New_York")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Timezone:", tzinfo.Timezone)
}

func getZone(zone string) (ret TimezoneInfo, err error) {

	url := fmt.Sprintf("https://api.ipgeolocation.io/timezone?apiKey=%s&tz=%s", APIKEY, zone)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("FAILED in: http.NewRequest", err)
		return ret, err
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("FAILED in: client.Do", err)
		return ret, err
	} else if res.StatusCode != 200 {
		return ret, errors.New("the server returned a non-200 status code: " + res.Status)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("FAILED in: io.ReadAll:", err)
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		fmt.Println("FAILED in: json.Unmarshal:", err)
		return ret, err
	}

	fmt.Println("Done without error")
	return ret, err
}
