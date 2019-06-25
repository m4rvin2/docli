package docstring

import "github.com/celicoo/docli/internal/reger"

// Parse returns the AST of the given doc string.
func Parse(s string) (d docstring) {
	if err := reger.Build(&d).ParseString(s, &d); err != nil {
		panic(err)
	}
	return
}
