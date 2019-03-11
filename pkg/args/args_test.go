package args_test

import (
	"os"
	"reflect"
	"strings"
	"testing"

	. "github.com/celicoo/docli/pkg/args"
	"github.com/celicoo/docli/pkg/text"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

func TestParse(t *testing.T) {
	var tests = []struct {
		args     []string
		expected Args
	}{
		{
			[]string{os.Args[0], "clone", "--depth=0", "--local"},
			Args{
				Arguments: []Argument{
					Argument{
						Identifier: "clone",
					},
					Argument{
						Identifier: "--depth",
						Value: Value{
							Assignment: "=",
							String:     "0",
						},
					},
					Argument{
						Identifier: "--local",
					},
				},
			},
		},
		{
			[]string{os.Args[0], "deploy", "--stage=production", "--region", "--verbose"},
			Args{
				Arguments: []Argument{
					Argument{
						Identifier: "deploy",
					},
					Argument{
						Identifier: "--stage",
						Value: Value{
							Assignment: "=",
							String:     "production",
						},
					},
					Argument{
						Identifier: "--region",
					},
					Argument{
						Identifier: "--verbose",
					},
				},
			},
		},
	}

	t.Log("Given the need to test the parser.")
	{
		for _, tt := range tests {
			t.Logf("\tWhen checking the AST produced by the \"[%s]\" arguments", strings.Join(tt.args, " "))
			{
				os.Args = tt.args
				args, _ := Parse(text.Text{})
				if reflect.DeepEqual(args, tt.expected) {
					t.Log("\t\tShould match the hard coded AST", checkMark)
				} else {
					t.Error("\t\tShould match the hard coded AST", ballotX)
				}
			}
		}
	}
}
