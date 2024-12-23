package cfg

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// Log to the screen
var l = log.New(os.Stderr, "commands:", log.Ldate|log.Ltime|log.Lshortfile)

var App cli.App

// this will init but panic. We will in fact return to main because we used recover()
// in the myPanic() function.
func init() {
	l.Println("cfg:init():Starting init in cfg package")

	l.Println("cfg:init():You realize you have to do this?")
	myPanic()
	l.Println("cfg:init():Finished init in cfg package")
}

// You can only recover the panic in the same function.
// you defer in the function and it will return to the calling function
// and then you can recover the panic.
// Panic will exit the function but recover() will return the control to the calling function.
func myPanic() {
	defer func() {
		if r := recover(); r != nil {
			l.Println("cfg:myPanic():Recovered in cfg package:", r)
		}
	}()
	l.Println("gonna panic, bitch")
	panic("Panicing in myPanic()")
}
