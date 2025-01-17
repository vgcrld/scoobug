package tz

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
	// "Asia/Kolkata":     "IST",
	"Australia/Sydney": "SYD",
}

// API key for the timezone API
var APIKEY string

func init() {
	data, err := os.ReadFile(".apikey")
	if err != nil {
		log.Fatalf("Failed to read API key from file: %v", err)
	}
	APIKEY = string(data)
}

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
	TimezoneOffset        float32 `json:"timezone_offset"`
	TimezoneOffsetWithDst float32 `json:"timezone_offset_with_dst"`
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

// Error struct for the errors
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

func ConvertToCtmFormat() []string {
	var ret []string
	var short, curr string
	var val string

	// this is not complete. The stdoff and dstoff need to be calculated
	// curr should be the gmt offset during standard time, always e.g. EDT (GMT-05:00)
	// off should be the gmt offsute during daylight time, always e.g. EST (GMT-04:00)
	// if there is no dst, then the curr and off are the same
	// the api I am using is making the conversion odd.
	for _, zone := range ZonesQueried {
		short = zone.GetShortName()
		curr = zone.DateTimeYmd[len(zone.DateTimeYmd)-5:]
		curr = curr[:3] + ":" + curr[3:]
		if zone.DstExists {
			off, _ := ConvertHourOffset(zone.DstStart.Duration)
			smon, sday, stime := splitDateTime(zone.DstStart.DateTimeBefore)
			emon, eday, etime := splitDateTime(zone.DstEnd.DateTimeBefore)
			val = fmt.Sprintf("%s (GMT%s) FROM %s.%s %s:00 TO %s.%s %s:00 (GMT%f) <<< NOT CORRECT\n", short, curr, sday, smon, stime, eday, emon, etime, off)
		} else {
			val = fmt.Sprintf("%s (GMT%s)\n", short, curr)
		}
		ret = append(ret, val)
	}
	return ret
}

func ConvertHourOffset(offset string) (float32, error) {
	var sign float32 = 1
	if offset[0] == '-' {
		sign = -1
	}
	hours := offset[1 : len(offset)-1]
	var hourValue float32
	_, err := fmt.Sscanf(hours, "%f", &hourValue)
	if err != nil {
		return 0, err
	}
	return sign * hourValue, nil
}

func splitDateTime(dateTime string) (string, string, string) {
	mon := dateTime[5:7]
	day := dateTime[8:10]
	time := dateTime[16:]
	return mon, day, time
}

func QueryAll() error {
	for long := range ZonesToFetch {
		log.Printf("Querying timezone: %s\n", long)
		err := QueryZoneAPI(long)
		if err != nil {
			fmt.Println(err)
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
