package uniq

import (
	"flag"
	"fmt"
	"io"
	"path/filepath"
)

func ParseCLI(args []string, version string, stdin io.Reader, stdout io.Writer) (result Configuration, err error) {
	flags := flag.NewFlagSet(fmt.Sprintf("%s @ %s", filepath.Base(args[0]), version), flag.ContinueOnError)
	flags.Usage = func() {
		_, _ = fmt.Fprintf(flags.Output(), "Usage of %s:\n", flags.Name())
		_, _ = fmt.Fprintf(flags.Output(), "%s [args ...]\n", filepath.Base(args[0]))
		_, _ = fmt.Fprintln(flags.Output(), "More details here.")
		flags.PrintDefaults()
	}

	result.EmitUnique = true

	err = flags.Parse(args[1:])
	if err != nil {
		return result, err
	}
	result.Source = stdin
	result.Target = stdout
	return result, nil
}
