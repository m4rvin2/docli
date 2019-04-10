package text

import "github.com/celicoo/docli/internal/reger"

// Parse returns the AST of the given doc string.
func Parse(doc string) (t Text) {
	if err := reger.Build(&t).ParseString(doc, &t); err != nil {
		panic(err)
	}
	return
}
