package main

import (
	"flag"

	"bin"
)

var debug = flag.Bool("g", false, "toggle debug mode")

func main() {
	flag.Parse()
	err := bin.StartClient(*debug)
	if err != nil {
		panic(err)
	}
}
