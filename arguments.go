package docli

import "github.com/celicoo/docli/v3/pkg/arguments"

// Arguments is the AST of the command-line arguments, starting after the program name.
var Arguments = arguments.Parse()
