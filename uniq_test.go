package uniq

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mdw-go/testing/should"
)

func TestReadError(t *testing.T) {
	boink := errors.New("boink")
	config, err := ParseCLI(t.Name(), &ErringReader{err: boink}, io.Discard, t.Name())
	should.So(t, err, should.BeNil)
	err = Process(config)
	should.So(t, err, should.WrapError, boink)
}

type ErringReader struct{ err error }

func (e *ErringReader) Read(p []byte) (n int, err error) { return 0, e.err }

func TestWriteError(t *testing.T) {
	boink := errors.New("boink")
	config, err := ParseCLI(t.Name(), strings.NewReader("hello"), &ErringWriter{err: boink}, t.Name())
	should.So(t, err, should.BeNil)
	err = Process(config)
	should.So(t, err, should.WrapError, boink)
}

type ErringWriter struct{ err error }

func (e *ErringWriter) Write(p []byte) (n int, err error) { return 0, e.err }

func testdata(path string) io.Reader {
	content, err := os.ReadFile(filepath.Join("testdata", path))
	if err != nil {
		panic(err)
	}
	return bytes.NewReader(content)
}
func process(t *testing.T, path string, args ...string) (result []string) {
	args = append([]string{t.Name()}, args...)
	var output bytes.Buffer
	config, err := ParseCLI(t.Name(), testdata(path), &output, args...)
	if err != nil {
		panic(err)
	}
	err = Process(config)
	if err != nil {
		panic(err)
	}
	return strings.Split(output.String(), "\n")
}
func TestDefaults(t *testing.T) {
	should.So(t, process(t, "12345-input.txt"), should.Equal, []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"",
	})

	should.So(t, process(t, "abc-input.txt"), should.Equal, []string{
		"Aa bb",
		"aa bb",
		"",
		"aa bb cc",
		"AA BB cc",
		"",
		"", // I'm not sure why this is necessary, but the builtin uniq command produces it too so I guess I'm good.
	})

	should.So(t, process(t, "different-at-end-input.txt"), should.Equal, []string{
		"4",
		"1",
		"",
	})
}
