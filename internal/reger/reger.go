package reger

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
)

const (
	// maxLookAhead is the number of tokens that the parser can read ahead to
	// perform disambiguation.
	// The current number is an arbitrary choice.
	maxLookAhead = 100

	// tokens contains the valid syntax elements.
	tokens = `(?P<Let>\p{L})` +
		`|(?P<Num>\p{N})` +
		`|(?P<Pun>\p{P})` +
		`|(?P<Sym>\p{S})` +
		`|(?P<Sep>\p{Z})` +
		`|(?P<Oth>\p{C})`
)

// Build constructs a parser for the given grammar.
func Build(grammar interface{}) *participle.Parser {
	l := lexer.Must(lexer.Regexp(tokens))
	return participle.MustBuild(grammar, participle.Lexer(l), participle.UseLookahead(maxLookAhead))
}
