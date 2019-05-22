package text

import "github.com/iancoleman/strcase"

// Text is both used as the grammar and the tree representation of the abstract
// syntactic structure of a doc string.
type Text struct {
	Arguments []Argument `(@@|`
	_         string     `Let|Num|Pun|Sym|Sep|Oth)*`
}

type Argument struct {
	Indentation string       `@('\n' (' '|'\t')+)`
	Identifiers []Identifier `@@+`
}

type Identifier struct {
	Name      string `@('-'|Let|Num)+`
	Separator string `@(',' ' ')?`
}

func (t *Text) argument(name string) Argument {
	for _, argument := range t.Arguments {
		for _, identifier := range argument.Identifiers {
			if identifier.Name == name {
				return argument
			}
		}
	}
	return Argument{}
}

// Identifiers returns all identifiers of an argument if any of them match the
// given value.
func (t *Text) Identifiers(name string) (identifiers []Identifier) {
	argument := t.argument(name)
	for _, identifier := range argument.Identifiers {
		identifiers = append(identifiers, identifier)
	}
	return
}

// NameAsCamelCase returns the Identifier's name as camel case.
// TODO(celicoo): improve this.
func (i *Identifier) NameAsCamelCase() string {
	return strcase.ToCamel(i.Name)
}
