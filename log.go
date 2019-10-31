package bin

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var debugLog = log.New(ioutil.Discard, "[DEBUG] ", 0)

func debugMode(debug bool) {
	if debug {
		debugLog.SetOutput(os.Stdout)
	}
	fmt.Printf("[DEBUG MODE] %t\n", debug)
}
