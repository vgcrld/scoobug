package other

import (
	"log"
	"os"
)

var l = log.New(os.Stderr, "other:", log.Ldate|log.Ltime|log.Lshortfile)

func init() {
	l.Println("Welcome to other...")
}
