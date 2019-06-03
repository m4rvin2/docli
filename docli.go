package docli

import (
	"os"
	"strings"
)

// Args holds the command-line arguments, starting after the program name.
var Args args

func init() {
	// Concatenate the command-line arguments using the U+001F character as
	// separator.
	s := strings.Join(os.Args[1:], "\u0009")
	Args.feed(s)
}
