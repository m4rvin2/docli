package docstring

import "github.com/celicoo/reger"

var tokens = reger.Tokens{
	"Let": `\p{L}`,
	"Num": `\p{N}`,
	"Pun": `\p{P}`,
	"Sym": `\p{S}`,
	"Sep": `\p{Z}`,
	"Oth": `\p{C}`,
}

var lexer reger.Lexer

func init() {
	lexer = reger.NewLexer(tokens)
}
