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
	flags := flag.NewFlagSet(fmt.Sprintf("%s @ %s", filepath.Base(os.Args[0]), Version), flag.ExitOnError)
	flags.Usage = func() {
		_, _ = fmt.Fprintf(flags.Output(), "Usage of %s:\n", flags.Name())
		_, _ = fmt.Fprintf(flags.Output(), "%s [args ...]\n", filepath.Base(os.Args[0]))
		_, _ = fmt.Fprintln(flags.Output(), "More details here.")
		flags.PrintDefaults()
	}
	_ = flags.Parse(os.Args[1:])

	uniq(os.Stdin, os.Stdout)
}

func uniq(input io.Reader, output io.Writer) {
	var previous string
	for reader := bufio.NewReader(input); ; {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if line == previous {
			continue
		}
		_, _ = fmt.Fprint(output, line)
		previous = line
	}
}
