package main

import (
	"fmt"

	vgcrld "github.com/vgcrld/scoobug/cfg"
)

type base struct {
	a int
}

type val struct {
	base
	b string
}

func main() {

	version := vgcrld.Version
	buildDate := vgcrld.BuildDate
	commitHash := vgcrld.CommitHash
	fmt.Println("Version:", version, "\nBuild Date:", buildDate, "\nCommit Hash:", commitHash)

	ff := val{
		base: base{a: 1},
		b:    "2",
	}
	fmt.Println(ff.base.a, ff.b)
}
