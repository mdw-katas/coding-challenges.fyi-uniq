package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var Version = "dev"

func main() {
	config, err := parseFlags(os.Args, Version)
	if err != nil {
		log.Fatal(err)
	}
	err = uniq(config)
	if err != nil {
		log.Fatal(err)
	}
}

func parseFlags(args []string, version string) (result Configuration, err error) {
	flags := flag.NewFlagSet(fmt.Sprintf("%s @ %s", filepath.Base(args[0]), version), flag.ContinueOnError)
	flags.Usage = func() {
		_, _ = fmt.Fprintf(flags.Output(), "Usage of %s:\n", flags.Name())
		_, _ = fmt.Fprintf(flags.Output(), "%s [args ...]\n", filepath.Base(args[0]))
		_, _ = fmt.Fprintln(flags.Output(), "More details here.")
		flags.PrintDefaults()
	}
	err = flags.Parse(args[1:])
	if err != nil {
		return result, err
	}
	result.Source = os.Stdin
	result.Target = os.Stdout
	return result, nil
}

type Configuration struct {
	Source io.Reader
	Target io.Writer
}

func uniq(config Configuration) error {
	var previous string
	for reader := bufio.NewReader(config.Source); ; {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("read: %w", err)
		}
		if line == previous {
			continue
		}
		_, err = fmt.Fprint(config.Target, line)
		if err != nil {
			return fmt.Errorf("write: %w", err)
		}
		previous = line
	}
}
