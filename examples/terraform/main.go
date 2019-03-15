package main

import (
	"log"

	"github.com/celicoo/docli"
)

var doc = `
Usage: terraform [--version] [--help] <command> [args]

The available commands for execution are listed below.
The most common, useful commands are shown first, followed by
less common or more advanced commands. If you're just getting
started with Terraform, stick with the common commands. For the
other commands, please read the help and docs before usage.

Common commands:
    apply              Builds or changes infrastructure
    console            Interactive console for Terraform interpolations
    destroy            Destroy Terraform-managed infrastructure
    fmt                Rewrites config files to canonical format
    get                Download and install modules for the configuration
    graph              Create a visual graph of Terraform resources
    import             Import existing infrastructure into Terraform
    init               Initialize a new or existing Terraform configuration
    output             Read an output from a state file
    plan               Generate and show an execution plan
    providers          Prints a tree of the providers used in the configuration
    push               Upload this Terraform module to Terraform Enterprise to run
    refresh            Update local state file against real resources
    show               Inspect Terraform state or plan
    taint              Manually mark a resource for recreation
    untaint            Manually unmark a resource as tainted
    validate           Validates the Terraform files
    version            Prints the Terraform version
    workspace          Workspace management

All other commands:
    debug              Debug output management (experimental)
    force-unlock       Manually unlock the terraform state
    state              Advanced state management
`

type Terraform struct {
	Commands
	AllOtherCommands
}

type Commands struct {
	Apply,
	Console,
	Destroy,
	Fmt,
	Get,
	Graph,
	Import,
	Init,
	Output,
	Plan,
	Providers,
	Push,
	Refresh,
	Show,
	Taint,
	Untaint,
	Validate,
	Version,
	Workspace bool
}

type AllOtherCommands struct {
	Debug,
	ForceUnlock,
	State bool
}

func main() {
	args, err := docli.Parse(doc)
	if err != nil {
		log.Fatal(err)
	}
	var terraform Terraform
	args.Bind(&terraform)
}
