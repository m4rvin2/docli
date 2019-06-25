package docli

import (
	"os"
	"strings"

	"github.com/celicoo/docli/internal/reger"
)

// Args returns the command-line arguments, starting after the program name.
func Args() (a args) {
	// Concatenate the command-line arguments using the U+001F character as
	// separator.
	s := strings.Join(os.Args[1:], "\u0009")
	reger.Build(&a).ParseString(s, &a)
	return
}
