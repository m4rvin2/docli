package docli_test

import (
	"fmt"
	"os"
	"strings"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	. "github.com/celicoo/docli"
)

var args interface{}

func theArguments(d *gherkin.DocString) error {
	os.Args = append(os.Args[:1], strings.Split(d.Content, " ")...)
	return nil
}

func iCallTheParser() error {
	args = Args()
	return nil
}

func theReturningValueShouldBeEqual(d *gherkin.DocString) error {
	if d.Content != fmt.Sprint(args) {
		return fmt.Errorf("\nExpected: \"%s\" \nActual: \"%s\"", d.Content, args)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^the arguments:$`, theArguments)
	s.Step(`^I call the parser$`, iCallTheParser)
	s.Step(`^the returning value should be equal$`, theReturningValueShouldBeEqual)
}
