package main

import (
	"log"

	"github.com/celicoo/docli"
)

var doc = `
Usage: appex [command]

Available Commands:
    alias        Create or update alias on functions
    build        Build a function
    delete       Delete functions
    deploy       Deploy functions and config
    docs         Output documentation
    exec         Command execution passthrough
    infra        Infrastructure management
    init         Initialize a project
    invoke       Invoke functions
    list         Output functions list
    logs         Output function logs
    metrics      Output function metrics
    rollback     Rollback functions
    upgrade      Upgrade apex to the latest stable release
    version      Print version of Apex

Flags:
    -C, --chdir string       Working directory
    -D, --dry-run            Perform a dry-run
        --endpoint string    AWS endpoint
    -e, --env string         Environment name
    -h, --help               help for apex
    -i, --iamrole string     AWS iamrole
    -l, --log-level string   Log severity level (default "info")
    -p, --profile string     AWS profile
    -r, --region string      AWS region

Use "apex [command] --help" for more information about a command.
`

type Apex struct {
	AvailableCommands
	Flags
}

type AvailableCommands struct {
	Alias,
	Build,
	Delete,
	Deploy,
	Docs,
	Exec,
	Infra,
	Init,
	Invoke,
	List,
	Logs,
	Metrics,
	Rollback,
	Upgrade,
	Version bool
}

type Flags struct {
	DryRun,
	Help bool
	Chdir,
	Endpoint,
	Env,
	Iamrole,
	LogLevel,
	Profile,
	Region string
}

func main() {
	args, err := docli.Parse(doc)
	if err != nil {
		log.Fatal(err)
	}
	var apex Apex
	args.Bind(&apex)
}
