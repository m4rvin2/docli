package docstring

import "github.com/celicoo/docli/v3/internal"

// Parse returns the AST of the docstring.
func Parse(docstring string) AST {
	a, e := internal.Parse(Grammar{}, docstring)
	if e != nil {
		internal.ReportError(e)
	}
	return a.(AST)
}

