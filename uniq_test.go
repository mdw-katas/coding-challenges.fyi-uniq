package uniq

import (
	"bytes"
	"io"
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
	return strings.Split(strings.TrimSpace(output.String()), "\n")
}
func TestUniqDefaults(t *testing.T) {
	should.So(t, uniq(t, "12345-input.txt"), should.Equal, []string{
		"1",
		"2",
		"3",
		"4",
		"5",
	})
}
