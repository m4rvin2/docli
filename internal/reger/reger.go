package reger

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
)

// maxLookAhead is the number of tokens that the parser can read ahead to
// perform disambiguation.
// The current number is an arbitrary choice.
const maxLookAhead = 100

// Parser is an alias of participle.Parser.
type Parser = *participle.Parser

// tokens contains the valid syntax elements.
var tokens = `(?P<Let>\p{L})` +
	`|(?P<Num>\p{N})` +
	`|(?P<Pun>\p{P})` +
	`|(?P<Sym>\p{S})` +
	`|(?P<Sep>\p{Z})` +
	`|(?P<Oth>\p{C})`

// Build constructs a parser for the given grammar.
func Build(grammar interface{}) Parser {
	l := lexer.Must(lexer.Regexp(tokens))
	return participle.MustBuild(grammar, participle.Lexer(l), participle.UseLookahead(maxLookAhead))
}
