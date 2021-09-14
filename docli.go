package docli

import "github.com/celicoo/docli/v2/pkg/args"

// Args returns the AST of the command-line arguments, starting after the
// program name.
func Args() args.Args {
	return args.Parse()
}
