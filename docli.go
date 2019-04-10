package docli

import "github.com/celicoo/docli-old/pkg/args"

// Args returns the AST of the command-line arguments.
func Args() args.Args {
	return args.Parse()
}
