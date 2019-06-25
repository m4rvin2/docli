package docli

import "fmt"

type InvalidArgumentError struct {
	Argument string
}

func (i *InvalidArgumentError) Error() string {
	return fmt.Sprintf("'%s' is not a valid argument", i.Argument)
}
