package docstring

import "github.com/celicoo/reger"

var parser reger.Parser

func init() {
	var d Docstring
	parser = reger.NewParser("docstring", d, lexer)
}
