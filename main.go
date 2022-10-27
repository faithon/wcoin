package main

import (
	"github.com/faithon/wcoin/cli"
	"github.com/faithon/wcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
