package arguments

type AST Grammar

// Bind binds the fields of the given struct with matching argument values.
func (a *AST) Bind(c command) {
	c.Run()

// INFO(celicoo): old implementation of bind.
//	v := reflect.ValueOf(c)
//	// Validate that given parameter is a pointer.
//	if v.Kind() != reflect.Ptr {
//		s := v.Type().Name()
//		e := fmt.Errorf("cannot use c (type %[1]v) as type *%[1]v in argument to Bind", s)
//		internal.ReportError(e)
//	}
//	e := v.Elem()
//	s := c.Doc()
//	d := docstring.Parse(s)
//arguments:
//	for i := range a.Arguments { // slight faster.
//		n := a.Arguments[i].Identifier
//		if n == "help" {
//			c.Help()
//			return
//		}
//		for _, identifier := range d.Identifiers(n) {
//			s := identifier.NameAsCamelCase()
//			f := e.FieldByName(s)
//			// Match found.
//			if f.CanSet() {
//				// Check if the field implements the Command interface.
//				c, ok := f.Addr().Interface().(command)
//				if ok {
//					// If it does implement the Command interface, delete the current
//					// argument from a.Arguments to prevent returning an
//					// ErrInvalidArgument by mistake.
//					a.Arguments = append(a.Arguments[:i], a.Arguments[i+1:]...)
//					a.Bind(c)
//					return
//				}
//				k := f.Kind()
//				if k == reflect.Bool {
//					f.SetBool(true)
//				} else {
//					s := a.Arguments[i].Value.String
//					// If the type of the field is string, use SetString to prevent losing
//					// information that otherwise could happen with fmt.Sscan.
//					if k == reflect.String {
//						f.SetString(s)
//					} else {
//						i := f.Addr().Interface()
//						fmt.Sscan(s, i)
//					}
//				}
//			} else {
//				continue
//			}
//			continue arguments
//		}
//		e := &InvalidArgumentError{n}
//		c.Error(e)
//		return
//	}
}
