package uniq

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type Configuration struct {
	EmitCounts      bool
	EmitUnique      bool // ignored if either EmitRepeated or EmitAllRepeated is set
	EmitRepeated    bool // ignored if EmitAllRepeated is set
	EmitAllRepeated bool

	SkipFields int
	SkipChars  int

	IgnoreCase bool

	Source io.Reader
	Target io.Writer
}

func Uniq(config Configuration) error {
	var previousLine bytes.Buffer
	for reader := bufio.NewReader(config.Source); ; {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("read: %w", err)
		}
		if bytes.Equal(line, previousLine.Bytes()) {
			continue
		}
		_, err = config.Target.Write(line)
		if err != nil {
			return fmt.Errorf("write: %w", err)
		}
		previousLine.Reset()
		_, _ = previousLine.Write(line)
	}
}
