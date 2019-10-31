package main

import (
	"flag"

	"bin"
)

var debug = flag.Bool("g", false, "toggle debug mode")

func main() {
	flag.Parse()
	err := bin.StartServer(*debug)
	if err != nil {
		panic(err)
	}
}
