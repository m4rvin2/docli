package text_test

import (
	"reflect"
	"testing"

	. "github.com/celicoo/docli/pkg/text"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

func TestParse(t *testing.T) {
	var tests = []struct {
		label    string
		doc      string
		expected Text
	}{
		{
			"docker",
			`
Usage:	docker [OPTIONS] COMMAND

A self-sufficient runtime for containers

Options:
      --config string      Location of client config files (default "/Users/celicoo/.docker")
  -D, --debug              Enable debug mode
  -H, --host list          Daemon socket(s) to connect to
  -l, --log-level string   Set the logging level ("debug"|"info"|"warn"|"error"|"fatal") (default "info")
      --tls                Use TLS; implied by --tlsverify
      --tlscacert string   Trust certs signed only by this CA (default "/Users/celicoo/.docker/ca.pem")
      --tlscert string     Path to TLS certificate file (default "/Users/celicoo/.docker/cert.pem")
      --tlskey string      Path to TLS key file (default "/Users/celicoo/.docker/key.pem")
      --tlsverify          Use TLS and verify the remote
  -v, --version            Print version information and quit

Management Commands:
  config      Manage Docker configs
  container   Manage containers
  image       Manage images
  network     Manage networks
  node        Manage Swarm nodes
  plugin      Manage plugins
  secret      Manage Docker secrets
  service     Manage services
  stack       Manage Docker stacks
  swarm       Manage Swarm
  system      Manage Docker
  trust       Manage trust on Docker images
  volume      Manage volumes

Commands:
  attach      Attach local standard input, output, and error streams to a running container
  build       Build an image from a Dockerfile
  commit      Create a new image from a container's changes
  cp          Copy files/folders between a container and the local filesystem
  create      Create a new container
  deploy      Deploy a new stack or update an existing stack
  diff        Inspect changes to files or directories on a container's filesystem
  events      Get real time events from the server
  exec        Run a command in a running container
  export      Export a container's filesystem as a tar archive
  history     Show the history of an image
  images      List images
  import      Import the contents from a tarball to create a filesystem image
  info        Display system-wide information
  inspect     Return low-level information on Docker objects
  kill        Kill one or more running containers
  load        Load an image from a tar archive or STDIN
  login       Log in to a Docker registry
  logout      Log out from a Docker registry
  logs        Fetch the logs of a container
  pause       Pause all processes within one or more containers
  port        List port mappings or a specific mapping for the container
  ps          List containers
  pull        Pull an image or a repository from a registry
  push        Push an image or a repository to a registry
  rename      Rename a container
  restart     Restart one or more containers
  rm          Remove one or more containers
  rmi         Remove one or more images
  run         Run a command in a new container
  save        Save one or more images to a tar archive (streamed to STDOUT by default)
  search      Search the Docker Hub for images
  start       Start one or more stopped containers
  stats       Display a live stream of container(s) resource usage statistics
  stop        Stop one or more running containers
  tag         Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE
  top         Display the running processes of a container
  unpause     Unpause all processes within one or more containers
  update      Update configuration of one or more containers
  version     Show the Docker version information
  wait        Block until one or more containers stop, then print their exit codes

Run 'docker COMMAND --help' for more information on a command.
				`,
			Text{
				Arguments: []Argument{
					Argument{
						Indentation: "\n      ",
						Identifiers: []Identifier{
							Identifier{
								Name: "--config",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name:      "-D",
								Separator: ", ",
							},
							Identifier{
								Name: "--debug",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name:      "-H",
								Separator: ", ",
							},
							Identifier{
								Name: "--host",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name:      "-l",
								Separator: ", ",
							},
							Identifier{
								Name: "--log-level",
							},
						},
					},
					Argument{
						Indentation: "\n      ",
						Identifiers: []Identifier{
							Identifier{
								Name: "--tls",
							},
						},
					},
					Argument{
						Indentation: "\n      ",
						Identifiers: []Identifier{
							Identifier{
								Name: "--tlscacert",
							},
						},
					},
					Argument{
						Indentation: "\n      ",
						Identifiers: []Identifier{
							Identifier{
								Name: "--tlscert",
							},
						},
					},
					Argument{
						Indentation: "\n      ",
						Identifiers: []Identifier{
							Identifier{
								Name: "--tlskey",
							},
						},
					},
					Argument{
						Indentation: "\n      ",
						Identifiers: []Identifier{
							Identifier{
								Name: "--tlsverify",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name:      "-v",
								Separator: ", ",
							},
							Identifier{
								Name: "--version",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "config",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "container",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "image",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "network",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "node",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "plugin",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "secret",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "service",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "stack",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "swarm",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "system",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "trust",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "volume",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "attach",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "build",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "commit",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "cp",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "create",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "deploy",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "diff",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "events",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "exec",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "export",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "history",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "images",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "import",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "info",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "inspect",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "kill",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "load",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "login",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "logout",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "logs",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "pause",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "port",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "ps",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "pull",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "push",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "rename",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "restart",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "rm",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "rmi",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "run",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "save",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "search",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "start",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "stats",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "stop",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "tag",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "top",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "unpause",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "update",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "version",
							},
						},
					},
					Argument{
						Indentation: "\n  ",
						Identifiers: []Identifier{
							Identifier{
								Name: "wait",
							},
						},
					},
				},
			},
		},
		{
			"git",
			`
usage: git [--version] [--help] [-C <path>] [-c <name>=<value>]
		   [--exec-path[=<path>]] [--html-path] [--man-path] [--info-path]
		   [-p | --paginate | --no-pager] [--no-replace-objects] [--bare]
		   [--git-dir=<path>] [--work-tree=<path>] [--namespace=<name>]
		   <command> [<args>]

These are common Git commands used in various situations:

start a working area (see also: git help tutorial)
   clone      Clone a repository into a new directory
   init       Create an empty Git repository or reinitialize an existing one

work on the current change (see also: git help everyday)
   add        Add file contents to the index
   mv         Move or rename a file, a directory, or a symlink
   reset      Reset current HEAD to the specified state
   rm         Remove files from the working tree and from the index

examine the history and state (see also: git help revisions)
   bisect     Use binary search to find the commit that introduced a bug
   grep       Print lines matching a pattern
   log        Show commit logs
   show       Show various types of objects
   status     Show the working tree status

grow, mark and tweak your common history
   branch     List, create, or delete branches
   checkout   Switch branches or restore working tree files
   commit     Record changes to the repository
   diff       Show changes between commits, commit and working tree, etc
   merge      Join two or more development histories together
   rebase     Reapply commits on top of another base tip
   tag        Create, list, delete or verify a tag object signed with GPG

collaborate (see also: git help workflows)
   fetch      Download objects and refs from another repository
   pull       Fetch from and integrate with another repository or a local branch
   push       Update remote refs along with associated objects

'git help -a' and 'git help -g' list available subcommands and some
concept guides. See 'git help <command>' or 'git help <concept>'
to read about a specific subcommand or concept.
					`,
			Text{
				Arguments: []Argument{
					Argument{
						Indentation: "\n   ",
						Identifiers: []Identifier{
							Identifier{
								Name: "clone",
							},
						},
					},
					Argument{
						Indentation: "\n   ",
						Identifiers: []Identifier{
							Identifier{
								Name: "init",
							},
						},
					},
					Argument{
						Indentation: "\n   ",
						Identifiers: []Identifier{
							Identifier{
								Name: "add",
							},
						},
					},
					Argument{
						Indentation: "\n   ",
						Identifiers: []Identifier{
							Identifier{
								Name: "mv",
							},
						},
					},
					Argument{
						Indentation: "\n   ",
						Identifiers: []Identifier{
							Identifier{
								Name: "reset",
							},
						},
					},
					Argument{
						Indentation: "\n   ",
						Identifiers: []Identifier{
							Identifier{
								Name: "rm",
							},
						},
					},
					Argument{
						Indentation: "\n   ",
						Identifiers: []Identifier{
							Identifier{
								Name: "bisect",
							},
						},
					},
					Argument{
						Indentation: "\n   ",
						Identifiers: []Identifier{
							Identifier{
								Name: "grep",
							},
						},
					},
					Argument{
						Indentation: "\n   ",
						Identifiers: []Identifier{
							Identifier{
								Name: "log",
							},
						},
					},
					Argument{
						Indentation: "\n   ",
						Identifiers: []Identifier{
							Identifier{
								Name: "show",
							},
						},
					},
					Argument{
						Indentation: "\n   ",
						Identifiers: []Identifier{
							Identifier{
								Name: "status",
							},
						},
					},
					Argument{
						Indentation: "\n   ",
						Identifiers: []Identifier{
							Identifier{
								Name: "branch",
							},
						},
					},
					Argument{
						Indentation: "\n   ",
						Identifiers: []Identifier{
							Identifier{
								Name: "checkout",
							},
						},
					},
					Argument{
						Indentation: "\n   ",
						Identifiers: []Identifier{
							Identifier{
								Name: "commit",
							},
						},
					},
					Argument{
						Indentation: "\n   ",
						Identifiers: []Identifier{
							Identifier{
								Name: "diff",
							},
						},
					},
					Argument{
						Indentation: "\n   ",
						Identifiers: []Identifier{
							Identifier{
								Name: "merge",
							},
						},
					},
					Argument{
						Indentation: "\n   ",
						Identifiers: []Identifier{
							Identifier{
								Name: "rebase",
							},
						},
					},
					Argument{
						Indentation: "\n   ",
						Identifiers: []Identifier{
							Identifier{
								Name: "tag",
							},
						},
					},
					Argument{
						Indentation: "\n   ",
						Identifiers: []Identifier{
							Identifier{
								Name: "fetch",
							},
						},
					},
					Argument{
						Indentation: "\n   ",
						Identifiers: []Identifier{
							Identifier{
								Name: "pull",
							},
						},
					},
					Argument{
						Indentation: "\n   ",
						Identifiers: []Identifier{
							Identifier{
								Name: "push",
							},
						},
					},
				},
			},
		},
		{
			"terraform",
			`
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
					`,
			Text{
				Arguments: []Argument{
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "apply",
							},
						},
					},
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "console",
							},
						},
					},
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "destroy",
							},
						},
					},
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "fmt",
							},
						},
					},
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "get",
							},
						},
					},
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "graph",
							},
						},
					},
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "import",
							},
						},
					},
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "init",
							},
						},
					},
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "output",
							},
						},
					},
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "plan",
							},
						},
					},
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "providers",
							},
						},
					},
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "push",
							},
						},
					},
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "refresh",
							},
						},
					},
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "show",
							},
						},
					},
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "taint",
							},
						},
					},
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "untaint",
							},
						},
					},
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "validate",
							},
						},
					},
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "version",
							},
						},
					},
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "workspace",
							},
						},
					},
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "debug",
							},
						},
					},
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "force-unlock",
							},
						},
					},
					Argument{
						Indentation: "\n\t",
						Identifiers: []Identifier{
							Identifier{
								Name: "state",
							},
						},
					},
				},
			},
		},
	}

	t.Log("Given the need to test the parser.")
	{
		for _, tt := range tests {
			t.Logf("\tWhen checking the AST produced by the \"%s\" help message", tt.label)
			{
				text, _ := Parse(tt.doc)
				if reflect.DeepEqual(text, tt.expected) {
					t.Log("\t\tShould match the hard coded AST", checkMark)
				} else {
					t.Error("\t\tShould match the hard coded AST", ballotX)
				}
			}
		}
	}
}
