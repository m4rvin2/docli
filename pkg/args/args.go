package args

import (
	"os"
	"strings"
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
	parser.Parse(args(), &a)
	return a
}
