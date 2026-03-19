package main

import (
	"log"

	"github.com/lukas-sgx/corel/cmd"
)

func main() {
	err := cmd.Cli()
	if err != nil {
		log.Fatal(err)
	}
}