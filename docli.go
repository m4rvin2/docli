package docli

import (
	"github.com/celicoo/docli/pkg/args"
	"github.com/celicoo/docli/pkg/text"
)

// Parse parses the command-line arguments based on the interface described in
// doc string.
func Parse(doc string) (args.Args, error) {
	t, err := text.Parse(doc)
	if err != nil {
		return args.Args{}, err
	}
	return args.Parse(t)
}
