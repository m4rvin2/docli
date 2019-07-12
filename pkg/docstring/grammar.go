package docstring

import "github.com/iancoleman/strcase"

// Docstring is both used as the grammar and the tree representation of the abstract
// syntactic structure of a doc string.
type Docstring struct {
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

func (d *Docstring) argument(name string) Argument {
	for _, a := range d.Arguments {
		for _, i := range a.Identifiers {
			if i.Name == name {
				return a
			}
		}
	}
	return Argument{}
}

// Identifiers returns all identifiers of an argument if any of them match the
// given value.
func (d *Docstring) Identifiers(name string) (identifiers []Identifier) {
	a := d.argument(name)
	for _, i := range a.Identifiers {
		identifiers = append(identifiers, i)
	}
	return
}

// NameAsCamelCase returns the identifier's name as camel case.
// TODO(celicoo): improve this.
func (i *Identifier) NameAsCamelCase() string {
	return strcase.ToCamel(i.Name)
}
