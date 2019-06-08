package docli

import (
	"fmt"
	"reflect"

	"github.com/celicoo/docli/internal/docstring"
)

// args is both used as the grammar and the tree representation of the abstract
// syntactic structure of the command-line arguments.
type args struct {
	Arguments []argument `(@@|`
	_         string     `Pun|Sym|Oth|'\u0009')*`
}

type argument struct {
	Identifier string `@('-'|Let|Num)+`
	Value      value  `@@?`
}

type value struct {
	Assignment string `@'='`
	String     string `@(Let|Num|Pun|Sym|Sep)+`
}

// Bind binds the fields of the given struct with matching argument values.
func (a *args) Bind(c command) {
	v := reflect.ValueOf(c)
	// Validate that given parameter is a pointer.
	if v.Kind() != reflect.Ptr {
		err := fmt.Errorf("cannot use c (type %[1]v) as type *%[1]v in argument to Bind", v.Type().Name())
		panic(err)
	}
	e := v.Elem()
	d := docstring.Parse(c.Doc())
arguments:
	for i := range a.Arguments { // slight faster.
		n := a.Arguments[i].Identifier
		if n == "help" {
			c.Help()
			return
		}
		for _, identifier := range d.Identifiers(n) {
			f := e.FieldByName(identifier.NameAsCamelCase())
			// Match found.
			if f.CanSet() {
				// Check if the field implements the Command interface.
				t, ok := f.Addr().Interface().(command)
				if ok {
					// If it does implement the Command interface, delete the current
					// argument from a.Arguments to prevent returning an
					// ErrInvalidArgument by mistake.
					a.Arguments = append(a.Arguments[:i], a.Arguments[i+1:]...)
					a.Bind(t)
					return
				}
				if f.Kind() == reflect.Bool {
					f.SetBool(true)
				} else {
					fmt.Sscan(a.Arguments[i].Value.String, f.Addr().Interface())
				}
			} else {
				continue
			}
			continue arguments
		}
		err := &InvalidArgumentError{n}
		c.Error(err)
		return
	}
	c.Run()
}
