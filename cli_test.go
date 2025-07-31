package uniq

import (
	"os"
	"testing"

	"github.com/mdw-go/testing/should"
)

func TestParseCLISuite(t *testing.T) {
	should.Run(&ParseCLISuite{T: should.New(t)}, should.Options.UnitTests())
}

type ParseCLISuite struct {
	*should.T
	args []string
}

func (this *ParseCLISuite) Setup() {
	this.AppendArgs("ccuniq")
}
func (this *ParseCLISuite) AppendArgs(args ...string) {
	this.args = append(this.args, args...)
}

func (this *ParseCLISuite) TestDefaults() {
	config, err := ParseCLI(this.args, "version", os.Stdin, os.Stdout)
	this.So(err, should.BeNil)
	this.So(config, should.Equal, Configuration{
		EmitUnique: true,
		Source:     os.Stdin,
		Target:     os.Stdout,
	})
}
