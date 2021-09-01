package args

type command interface {
	// Doc returns the docstring of a command.
	Doc() string
	// Run is called when the fields of a command are bound with matching
	// command-line arguments.
	Run()
	// Help is called when a user runs "<command> help".
	Help()
	// Error is called when unregistered command-line arguments are passed.
	Error(error)
}
