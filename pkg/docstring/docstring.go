package docstring

// Parse returns the AST of the given doc string.
func Parse(s string) Docstring {
	var d Docstring
	parser.Parse(s, &d)
	return d
}
