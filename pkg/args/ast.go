package args

import (
	"fmt"
	"reflect"

	"github.com/celicoo/docli/pkg/text"
	"github.com/iancoleman/strcase"
)

// Args is both used as the grammar and the tree representation of the abstract
// syntactic structure of the command-line arguments.
type Args struct {
	text      text.Text
	Arguments []Argument `(@@|`
	_         string     `Pun|Sym|Oth|'\u0009')*`
}

type Argument struct {
	Identifier string `@('-'|Let|Num)+`
	Value      Value  `@@?`
}

type Value struct {
	Assignment string `@'='`
	String     string `@(Let|Num|Pun|Sym|Sep)+`
}

// validate validates the arguments given on the command-line.
func (a *Args) validate() error {
	for _, argument := range a.Arguments {
		if a.text.IsValidIdentifier(argument.Identifier) {
			continue
		}
		return fmt.Errorf("'%s' is not a valid argument", argument.Identifier)
	}
	return nil
}

// Bind binds the fields of the given struct with matching argument values.
func (a *Args) Bind(i interface{}) {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == reflect.Invalid || v.Kind() != reflect.Struct || v.NumField() == 0 {
		return
	}
arguments:
	for _, argument := range a.Arguments {
		for _, identifier := range a.text.Identifiers(argument.Identifier) {
			var retry bool
		retry:
			// strcamel doesn't support non-ASCII words unfortunately.
			// refact(celicoo)
			if retry {
				identifier = strcase.ToCamel(identifier)
			}
			f := v.FieldByName(identifier)
			// Match found.
			if f.CanSet() {
				if f.Kind() == reflect.Bool {
					f.SetBool(true)
				} else {
					fmt.Sscan(argument.Value.String, f.Addr().Interface())
				}
				continue arguments
			} else if retry {
				// Next identifier.
				continue
			}
			retry = true
			goto retry
		}
	}
}
