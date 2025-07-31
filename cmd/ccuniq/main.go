package main

import (
	"log"
	"os"

	"github.com/mdw-katas/coding-challenges.fyi-uniq"
)

var Version = "dev"

func main() {
	config, err := uniq.ParseCLI(os.Args, Version, os.Stdin, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
	err = uniq.Uniq(config)
	if err != nil {
		log.Fatal(err)
	}
}
