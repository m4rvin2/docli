package internal

import (
	"fmt"
	"os"
)

// ReportError prints err to stderr and exits the program with status code 1.
func ReportError(err error) {
	fmt.Fprint(os.Stderr, err)
	os.Exit(1)
}
