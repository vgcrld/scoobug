package main

import (
	"fmt"

	"github.com/vgcrld/scoobug/samples/timezones/tz"
)

func main() {

	p := fmt.Println

	tz.QueryAll()
	p("Fetched: ", len(tz.ZonesQueried))
	// p(tz.Errors)
	tz.ZonesQueried[0].PrintJSON()

}
