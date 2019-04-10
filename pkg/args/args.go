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

// Parse returns the AST of the command-line arguments.
func Parse() (a Args) {
	if err := reger.Build(&a).ParseString(args(), &a); err != nil {
		panic(err)
	}
	return
}
