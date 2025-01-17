package main

import (
	"fmt"
	"os"

	"github.com/vgcrld/scoobug/samples/timezones/tz"
)

func main() {

	if err := tz.QueryAll(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tz.PrintJSON()
	fmt.Println(tz.ConvertToCtmFormat())
}
