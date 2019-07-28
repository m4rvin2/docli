package args

import "github.com/celicoo/reger"

var parser reger.Parser

func init() {
	var a Args
	parser = reger.NewParser("args", a, lexer)
}
