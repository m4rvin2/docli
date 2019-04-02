package args

import (
	"os"
	"strings"

	"github.com/celicoo/docli/internal/reger"
	"github.com/celicoo/docli/pkg/args/internal/text"
)

// args returns the command-line arguments, starting after the program name.
func args() string {
	// Concatenate the command-line arguments using the U+001F character as
	// separator.
	return strings.Join(os.Args[1:], "\u0009")
}

// Parse returns the AST of the command-line arguments.
func Parse(t text.Text) (Args, error) {
	a, parser := Args{}, reger.Build(&Args{})
	if err := parser.ParseString(args(), &a); err != nil {
		return Args{}, err
	}
	a.text = t
	return a, a.validate()
}
