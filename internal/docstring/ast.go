package docstring

import "github.com/iancoleman/strcase"

// docstring is both used as the grammar and the tree representation of the abstract
// syntactic structure of a doc string.
type docstring struct {
	Arguments []argument `(@@|`
	_         string     `Let|Num|Pun|Sym|Sep|Oth)*`
}

type argument struct {
	Indentation string       `@('\n' (' '|'\t')+)`
	Identifiers []identifier `@@+`
}

type identifier struct {
	Name      string `@('-'|Let|Num)+`
	Separator string `@(',' ' ')?`
}

func (d *docstring) argument(name string) argument {
	for _, a := range d.Arguments {
		for _, i := range a.Identifiers {
			if i.Name == name {
				return a
			}
		}
	}
	return argument{}
}

// Identifiers returns all identifiers of an argument if any of them match the
// given value.
func (d *docstring) Identifiers(name string) (identifiers []identifier) {
	a := d.argument(name)
	for _, i := range a.Identifiers {
		identifiers = append(identifiers, i)
	}
	return
}

// NameAsCamelCase returns the identifier's name as camel case.
// TODO(celicoo): improve this.
func (i *identifier) NameAsCamelCase() string {
	return strcase.ToCamel(i.Name)
}
