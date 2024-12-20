package main

import (
	"fmt"

	"github.com/vgcrld/scoobug/cfg"
)

func main() {
	fmt.Printf("Hello, World! Date: %s, Version: %s\n", cfg.BuildDate, cfg.Version)
}
