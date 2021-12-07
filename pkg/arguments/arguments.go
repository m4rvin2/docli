package arguments

import (
	"os"
	"strings"

	"github.com/celicoo/docli/v3/internal"
)

// Parse returns the AST of the command-line arguments, starting after the
// program name.
func Parse() AST {
	a, e := internal.Parse(
		Grammar{},
		// Concatenate the command-line arguments using the U+001F character as
		// separator.
		strings.Join(os.Args[1:], "\u0009"),
	)
	if e != nil {
		internal.ReportError(e)
	}
	return a.(AST)
}
