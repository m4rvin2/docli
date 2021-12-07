package internal

import (
	"regexp"

	"github.com/celicoo/alligopher"
)

func init() {
	t := alligopher.TokenTypes{
		"LET": alligopher.TokenType{
			regexp.MustCompile(`\p{L}`).FindStringIndex,
			nil,
		},
		"NUM": alligopher.TokenType{
			regexp.MustCompile(`\p{N}`).FindStringIndex,
			nil,
		},
		"PUN": alligopher.TokenType{
			regexp.MustCompile(`\p{P}`).FindStringIndex,
			nil,
		},
		"SYM": alligopher.TokenType{
			regexp.MustCompile(`\p{S}`).FindStringIndex,
			nil,
		},
		"SEP": alligopher.TokenType{
			regexp.MustCompile(`\p{Z}`).FindStringIndex,
			nil,
		},
		"OTH": alligopher.TokenType{
			regexp.MustCompile(`\p{C}`).FindStringIndex,
			nil,
		},
	}
	alligopher.Lexer.SetTokenTypes(t)
}

func Parse(grammar interface{}, input string) (interface{}, error) {
	return nil, nil
	//return alligopher.
	//	NewParser(grammar).
	//	Parse(input)
}
