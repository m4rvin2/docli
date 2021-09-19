package args

import (
	"os"
	"strings"

	"github.com/celicoo/docli/v2/internal/reger"
)

// Parse returns the AST of the command-line arguments, starting after the
// program name.
func Parse() Args {
	// Concatenate the command-line arguments using the U+001F character as
	// separator.
	s := strings.Join(os.Args[1:], "\u0009")
	// Ignore any errors that might happen when parsing the command-line
	// arguments. Although this might seem odd, it's preferable over the bad
	// error handling of Participle.
	// INFO(celicoo): errors will be nicely handled once I'm done developing
	// a replacement for Participle.
	var a Args
	reger.Build(&a).ParseString(s, &a)
	return a
}
