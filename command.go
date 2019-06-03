package docli

type command interface {
	// Doc returns the doc string of the [command].
	Doc() string

	// Run is called if the bind succeeds.
	Run()

	// Help is called when a user calls help [command].
	Help()

	// Error is called when a user uses the [command] wrongly.
	Error(error)
}
