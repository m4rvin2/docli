package text

import "github.com/celicoo/docli/internal/reger"

// Parse returns the AST of the given text.
func Parse(doc string) (Text, error) {
	t, parser := Text{}, reger.Build(&Text{})
	return t, parser.ParseString(doc, &t)
}
