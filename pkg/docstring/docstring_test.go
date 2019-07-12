package docstring_test

import (
	"fmt"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	. "github.com/celicoo/docli/pkg/docstring"
)

var (
	doc       string
	docstring Docstring
)

func theDocstring(d *gherkin.DocString) error {
	doc = d.Content
	return nil
}

func iCallTheParser() error {
	docstring = Parse(doc)
	return nil
}

func theReturningValueShouldBeEqual(d *gherkin.DocString) error {
	if d.Content != fmt.Sprint(docstring) {
		return fmt.Errorf("\nExpected: \"%s\" \nActual: \"%s\"", d.Content, docstring)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^the docstring:$`, theDocstring)
	s.Step(`^I call the parser$`, iCallTheParser)
	s.Step(`^the returning value should be equal$`, theReturningValueShouldBeEqual)
}
