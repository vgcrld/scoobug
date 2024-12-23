package cfg

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// Log to the screen
var Log = log.New(os.Stderr, "cfg:", log.LstdFlags)

var App cli.App

// this will init but panic. We will in fact return to main because we used recover()
// in the myPanic() function.
func init() {
	Log.Println("cfg:init():Starting init in cfg package")

	Log.Println("cfg:init():You realize you have to do this?")
	myPanic()
	Log.Println("cfg:init():Finished init in cfg package")
}

// You can only recover the panic in the same function.
// you defer in the function and it will return to the calling function
// and then you can recover the panic.
// Panic will exit the function but recover() will return the control to the calling function.
func myPanic() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("cfg:myPanic():Recovered in cfg package:", r)
		}
	}()
	log.Println("cfg:myPanic():nope, gonna panic, bitch")
	panic("Panicing in myPanic()")
}
