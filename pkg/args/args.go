package args

import (
	"os"
	"strings"

	"github.com/celicoo/docli/internal/reger"
)

// args returns the command-line arguments, starting after the program name.
func args() string {
	// Concatenate the command-line arguments using the U+001F character as
	// separator.
	return strings.Join(os.Args[1:], "\u0009")
}

// Parse returns the AST of the command-line arguments, starting after the
// program name.
func Parse() Args {
	var a Args
	// We ignore any errors that might happen when parsing the command-line
	// arguments. Although this looks odd, it's preferable over the bad error
	// handling of Participle.
	reger.Build(&a).ParseString(args(), &a)
	return a
}
