package uniq

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mdw-go/testing/should"
)

func testdata(path string) io.Reader {
	content, err := os.ReadFile(filepath.Join("testdata", path))
	if err != nil {
		panic(err)
	}
	return bytes.NewReader(content)
}
func uniq(t *testing.T, path string, args ...string) (result []string) {
	args = append([]string{t.Name()}, args...)
	var output bytes.Buffer
	config, err := ParseCLI(t.Name(), testdata(path), &output, args...)
	if err != nil {
		panic(err)
	}
	err = Uniq(config)
	if err != nil {
		panic(err)
	}
	return strings.Split(output.String(), "\n")
}
func TestUniqReadError(t *testing.T) {
	boink := errors.New("boink")
	config, err := ParseCLI(t.Name(), &ErringReader{err: boink}, ioutil.Discard, t.Name())
	should.So(t, err, should.BeNil)
	err = Uniq(config)
	should.So(t, err, should.WrapError, boink)
}

type ErringReader struct{ err error }

func (e *ErringReader) Read(p []byte) (n int, err error) { return 0, e.err }

func TestUniqWriteError(t *testing.T) {
	boink := errors.New("boink")
	config, err := ParseCLI(t.Name(), strings.NewReader("hello"), &ErringWriter{err: boink}, t.Name())
	should.So(t, err, should.BeNil)
	err = Uniq(config)
	should.So(t, err, should.WrapError, boink)
}

type ErringWriter struct{ err error }

func (e *ErringWriter) Write(p []byte) (n int, err error) { return 0, e.err }

func TestUniqDefaults(t *testing.T) {
	should.So(t, uniq(t, "12345-input.txt"), should.Equal, []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"",
	})

	should.So(t, uniq(t, "abc-input.txt"), should.Equal, []string{
		"Aa bb",
		"aa bb",
		"",
		"aa bb cc",
		"AA BB cc",
		"",
		"", // I'm not sure why this is necessary, but the builtin uniq command produces it too so I guess I'm good.
	})

	should.So(t, uniq(t, "different-at-end-input.txt"), should.Equal, []string{
		"4",
		"1",
		"",
	})
}
