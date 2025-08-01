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

func Process(config Configuration) error {
	var previousLine bytes.Buffer
	reader := bufio.NewReader(config.Source)
	for {
		line, readErr := reader.ReadBytes('\n')
		if readErr != nil && readErr != io.EOF {
			return fmt.Errorf("read: %w", readErr)
		}
		if len(line) == 0 && readErr == io.EOF {
			return nil
		}
		line = bytes.TrimSuffix(line, []byte{'\n'})
		if bytes.Equal(line, previousLine.Bytes()) {
			continue
		}
		_, writeErr := config.Target.Write(append(line, '\n'))
		if writeErr != nil {
			return fmt.Errorf("write: %w", writeErr)
		}
		if readErr == io.EOF {
			return nil
		}
		previousLine.Reset()
		_, _ = previousLine.Write(line)
	}
}
