package cli

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/faithon/wcoin/explorer"
	"github.com/faithon/wcoin/rest"
)

func usage() {
	fmt.Println("Welcome to my blockchain")
	fmt.Println("Please use the following flags:\n")
	fmt.Println("-port 4000: 	Sets port")
	fmt.Println("-mode rest: 	Sets mode")
	runtime.Goexit()
}

func Start() {
	port := flag.Int("port", 4000, "Sets port of the server")
	mode := flag.String("mode", "rest", "Choose 1. html, 2. rest")

	flag.Parse()
	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
