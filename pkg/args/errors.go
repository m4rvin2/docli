package args

import "fmt"

type InvalidArgumentError struct {
	Argument string
}

func (e *InvalidArgumentError) Error() string {
	return fmt.Sprintf("'%s' is not a valid argument", e.Argument)
}
