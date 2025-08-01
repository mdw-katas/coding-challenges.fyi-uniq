package main

import (
	"log"
	"os"

	"github.com/mdw-katas/coding-challenges.fyi-uniq"
)

var Version = "dev"

func main() {
	config, err := uniq.ParseCLI(Version, os.Stdin, os.Stdout, os.Args...)
	if err != nil {
		log.Fatal(err)
	}
	err = uniq.Uniq(config)
	if err != nil {
		log.Fatal(err)
	}
}
