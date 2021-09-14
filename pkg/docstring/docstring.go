package docstring

import "github.com/celicoo/docli/v2/internal/reger"

// Parse returns the AST of the given doc string.
func Parse(s string) Docstring {
	var d Docstring
	if err := reger.Build(&d).ParseString(s, &d); err != nil {
		panic(err)
	}
	return d
}
