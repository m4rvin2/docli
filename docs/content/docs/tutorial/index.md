---
title: Tutorial
---

# Tutorial

Welcome to the Docli Tutorial! Here you'll be introduced to the key Docli concepts. If you get stuck at any point in 
the process, feel free to [download](https://github.com/celicoo/docli/tree/main/examples/git) a working example of the 
completed CLI app.

## Confirming that Docli is installed

Before starting the tutorial, let's make sure that you have Docli installed. Go ahead and create a `main.go` file and paste the following content inside:

{{<highlight go>}}
package main

import (
	"github.com/alecthomas/repr"
	"github.com/celicoo/docli/v2"
)

func main() {
  args := docli.Args()
  repr.Println(args)
}
{{</highlight>}}

Now build and run it:

{{<highlight text>}}
$ go build
$ ./main

docli.args{
}
{{</highlight>}}

**Note**: the output is the [Abstract Syntax Structure](https://en.wikipedia.org/wiki/Abstract_syntax_tree) of the command-line arguments, and it's empty because you didn't pass any arguments. If you run it again passing arguments you should see a different output. 

## Creating a new CLI app

### Directory structure

If you like [Cobra](http://cobra.dev), you'll feel at home with this. While you're welcome to provide your own organization, typically a Docli-based app will follow the following directory structure:

{{< highlight text >}}
▾ app/
  ▾ cmd/
    root.go
  main.go
{{< /highlight >}}

Go ahead and create the above structure, but instead of `app`, let's name the root directory to `git`.

#### **app/**

This is the root of our app. The `main.go` file goes into this directory and it's responsible for calling the function that initializes Docli:

{{<highlight go>}}
package main

import "<path>/git"

func main() {
	cmd.Execute()
}
{{</highlight>}}

**Note**: Make sure to replace the `<path>` placeholder.

#### **cmd/**

This is where the commands are stored. The `root.go` is responsible for the logic of the **root command** and the `Execute()` function:

{{<highlight go>}}
package cmd

const version = "0.0.1"

type Git struct {}

func (g *Git) Doc() string {
	return ""
}

func (g *Git) Run() {
}

func (g *Git) Help() {
}

func (g *Git) Error(err error) {
}

func Execute() {
}
{{</highlight>}}

The `Git` struct represents the root command in our CLI app and it must implement the [docli.command](/) interface.

### Writing the docstring

Like YAML or Python, the docstring is a line-oriented language that uses indentation to define structure. Lines beginning with either spaces or tabs are used to register arguments and commands (commands are explained a little later in the tutorial). These arguments and commands can have letters of any language, numbers, and dashes. 

Let’s go back to the root.go file and replace the [Doc()](/) method of our Git struct with:

{{<highlight go>}}
func (g *Git) Doc() string {
	return `usage: git <command>

commands:
  c, clone      clone a repository into a new directory

arguments:

  v, --version  print version
	
Use "git <command> help" for more information about the <command>.`
}
{{</highlight>}}

By convention, dashes are used in front of arguments, but not commands. You can use this convention with Docli, but it's not necessary. 

### Accessing command-line argument values

Now that we’ve registered the arguments and commands in the docstring, we’ll need to define them as fields in the `Git` struct, so we can access their values. Fields that represent arguments will be primitive types and those that represent commands will be user-defined types that implement the [docli.command](/) interface, just like the root command does.

Regardless of how many aliases your arguments or commands have , you'll only need to (and only should) define one field in the struct. We suggest using the longer identifier, but you can use whichever you’d like. 

In our example, we have one command and one argument, and each has two identifiers. Let’s add these to our `Git` struct:

{{<highlight go>}}
type Git struct {
	Clone   Clone
	Version bool
}
{{</highlight>}}

For every command, you’ll need to create a file to hold the type that represents the command. In our example, the Clone field represents the `clone` command,  so we need to create a file that holds this type. Let's go ahead and create the `clone.go` file inside the `cmd` directory and paste the following code inside:

{{<highlight go>}}
package cmd

import (
	"fmt"
	"log"

	"github.com/alecthomas/repr"
)

type Clone struct {
	Repository,
	Directory string
	Verbose,
	Quiet,
	Progress,
	NoCheckout,
	Bare,
	Mirror,
	Local,
	NoHardlinks,
	Shared bool
	RecurseSubmodules string
	Jobs              int
	Template,
	Reference,
	ReferenceIfAble string
	Dissociate bool
	Origin,
	Branch,
	UploadPack string
	Depth int
	ShallowSince,
	ShallowExclude string
	SingleBranch,
	NoTags,
	ShallowSubmodules bool
	SeparateGitDir,
	Config string
	Ipv4,
	Ipv6 bool
	Filter string
}

func (c *Clone) Doc() string {
	return `usage: git clone [<arguments>] --repository=<url> --directory=<directory>

arguments:

  -r, --repository                 repository to clone
  -d, --directory                  path to directory
  -v, --verbose                    be more verbose
  -q, --quiet                      be more quiet
  --progress                       force progress reporting
  -n, --no-checkout                don't create a checkout
  --bare                           create a bare repository
  --mirror                         create a mirror repository (implies bare)
  -l, --local                      to clone from a local repository
  --no-hardlinks                   don't use local hardlinks, always copy
  -s, --shared                     setup as shared repository
  --recurse-submodules=<pathspec>  initialize submodules in the clone
  -j, --jobs=<n>                   number of submodules cloned in parallel
  --template=<template-directory>  directory from which templates will be used
  --reference=<repo>               reference repository
  --reference-if-able=<repo>       reference repository
  --dissociate                     use --reference only while cloning
  -o, --origin=<name>              use <name> instead of 'origin' to track upstream
  -b, --branch=<branch>            checkout <branch> instead of the remote's HEAD
  -u, --upload-pack=<path>         path to git-upload-pack on the remote
  --depth=<depth>                  create a shallow clone of that depth
  --shallow-since=<time>           create a shallow clone since a specific time
  --shallow-exclude=<revision>     deepen history of shallow clone, excluding rev
  --single-branch                  clone only one branch, HEAD or --branch
  --no-tags                        don't clone any tags, and make later fetches not to follow them
  --shallow-submodules             any cloned submodules will be shallow
  --separate-git-dir=<gitdir>      separate git dir from working tree
  -c, --config=<key=value>         set config inside the new repository
  -4, --ipv4                       use IPv4 addresses only
  -6, --ipv6                       use IPv6 addresses only
  --filter=<args>                  object filtering
`
}

func (c *Clone) Run() {
	repr.Println(c)
}

func (c *Clone) Help() {
	fmt.Println(c.Doc())
}

func (c *Clone) Error(err error) {
	log.Fatal(err)
}
{{</highlight>}}

In this tutorial, we won’t go into detail about the logic of the Clone command because its functionality is pretty straightforward.

### Writing the Methods’ Logic

A specific method will run when a certain condition is met.

#### **Error method**

The Error method will run when the user passes an argument or command that is not registered in the docstring. Usually, you'll just print the error message to `stderr` and exit the program. In our example, we'll ignore invalid arguments and force the Run method to run, otherwise we'll print the error message to `stderr` and exit the program:

{{<highlight go>}}
func (g *Git) Error(err error) {
	switch err.(type) {
	case *args.InvalidArgumentError:
		// Ignore InvalidArgumentError.
		g.Run()
	default:
		log.Fatal(err)
	}
}
{{</highlight>}}

The `InvalidArgumentError` is within the `args` package, so make sure to add `"github.com/celicoo/docli/v2/args"` to 
the import declarations.

#### **Help method**

The Help method will run when the user passes the `help` argument. Usually, you'll just print the docstring to `stdout`; that’s what we’ll do in our example:

{{<highlight go>}}
func (g *Git) Help() {
	fmt.Println(g.Doc())
}
{{</highlight>}}

#### **Run method**

The Run method will run when no error is found and an argument other than `help` is called. In our example we'll print the version if the `--version` (or `-v`) command-line argument is passed, otherwise we'll just print "Hello, world!":

{{<highlight go>}}
func (g *Git) Run() {
	if g.Version {
		fmt.Println(version)
		return
	}
	fmt.Println("Hello, world!")
}
{{</highlight>}}

### Initializing Docli

We created a CLI app with one command and one argument. But if you try to execute it, it won’t run. There is one piece missing: initializing Docli.

{{<highlight go>}}
func Execute() {
	var g Git
	args := docli.Args()
	args.Bind(&g)
}
{{</highlight>}}

### Testing the CLI app

The very last step is to make sure it’s actually working. To do that, we need to build and run it:

{{<highlight text>}}
$ go build
$ ./git

Hello, world!
{{</highlight>}}

Unlike other packages, Docli doesn't allow values to be assigned to arguments without using the `=` operator. You need to use the `=` operator, otherwise the internal parser will think you’re passing an argument and, as a result, the `Error` method of the struct that represents the command that you're executing will be called. The `clone` command was added to help you understand this, so go ahead and play with it.

Congratulations! You’ve completed the tutorial. Thank you for reading, and I hope you enjoyed it.

If you have any question, please don’t hesitate to [open an issue](https://github.com/celicoo/docli/issues).
