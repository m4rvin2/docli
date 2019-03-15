# Tutorial

Once you've [installed]({{< relref "/docs/installation" >}}) Docli via `go get`, you'll be able to import it on your new project. Go ahead and create a file called `git.go` and paste the following content inside:

{{< highlight go >}}
package main

import (
	"log"

	"github.com/alecthomas/repr"
	"github.com/celicoo/docli"
)

var doc = ``

func main() {
	args, err := docli.Parse(doc)
	if err != nil {
		log.Fatal(err)
	}
	repr.Println(args)
}
{{< /highlight >}}

Let’s build and run it to make sure everything is working correctly:

{{< highlight text >}}
$ go build
$ ./git

args.Args{
  text: text.Text{
  },
}
{{< /highlight >}}

The output is the [Abstract Syntax Structure](https://en.wikipedia.org/wiki/Abstract_syntax_tree) of the doc string. Right now it’s empty because the doc string is empty, or in other words, the command-line interface doesn't exist.

## Write a command-line interface

We'll write a simplified version of the `git` CLI. To do that, the first step is to define the argument identifiers. The pattern is simple:

{{< highlight text >}}
<indent with spaces or tabs><identifier [, <identifier>...]>
{{< /highlight >}}

The identifiers can have letters of any language, numbers and dashes. Unlike other libraries, Docli doesn't enforce the `short-long` convention.

**Note:** anything that doesn't match the pattern above is ignored.

Now that you know how to define argument identifiers, let's go back to the `git.go` file and replace the `doc` variable of our command-line interface with:

{{< highlight go >}}
var doc = `
Simple git

Usage: git [command]

Commands:
     init         create an empty Git repository or reinitialize an existing one
  h, help, 助けて  help message
`
{{< /highlight >}}

By looking at our command-line interface, can you guess how many **arguments** does it have? Try to make a guess before moving forward.

If you thought **two** you got it! Let's build and rerun it to confirm:

{{< highlight text >}}
$ go build
$ ./git

args.Args{
  text: text.Text{
    Arguments: []text.Argument{
      text.Argument{
        Indentation: "\n     ",
        Identifiers: []text.Identifier{
          text.Identifier{
            Name: "init",
          },
        },
      },
      text.Argument{
        Indentation: "\n  ",
        Identifiers: []text.Identifier{
          text.Identifier{
            Name: "h",
            Separator: ", ",
          },
          text.Identifier{
            Name: "help",
            Separator: ", ",
          },
          text.Identifier{
            Name: "助けて",
          },
        },
      },
    },
  },
}
{{< /highlight >}}

The **h**, **help**, and **助けて** are identifiers of the same argument. If you are confused, please go back to the [Write a command-line interface](#write-a-command-line-interface) section.

## Bind struct fields to command-line arguments

Now that we’ve written our command-line interface and verified that Docli correctly parses it, it’s time to write a struct to be bind to the command-line arguments values. To do so, paste the following struct after the doc string:

{{< highlight go >}}
type Git struct {
	Init,
	Help bool
}
{{< /highlight >}}

**Note:** you can use **any** primitive type in your struct fields and use embedded structs to separate the categories of your arguments, you'll find examples of the latest in the [examples page](https://github.com/celicoo/docli/tree/master/examples).

Replace the call to `repr.Println(args)` with:

{{< highlight go >}}
var git Git
args.Bind(&git)
if git.Help {
    fmt.Println(doc)
}
{{< /highlight >}}

The final code should be similar to the following:

{{< highlight go >}}
package main

import (
	"log"

	"github.com/celicoo/docli"
)

var doc = `
Simple git

Usage: git [command]

Commands:
     init         create an empty Git repository or reinitialize an existing one
  h, help, 助けて  help message
`

type Git struct {
	Init,
	Help bool
}

func main() {
	args, err := docli.Parse(doc)
	if err != nil {
		log.Fatal(err)
	}
    var git Git
    args.Bind(&git)
    if git.Help {
        fmt.Println(doc)
    }
}
{{< /highlight >}}

If we now build and run it:

{{< highlight text >}}
$ go build
$ ./git 助けて

Simple git

Usage: git [command]

Commands:
     init         create an empty Git repository or reinitialize an existing one
  h, help, 助けて  help message
{{< /highlight >}}

Boolean arguments are special because they don't require explicitly assignment like arguments of other types.

For example, if we decide to add an argument `--author` of type `string`, the correct way to bind a value to it would be:

{{< highlight text >}}
$ ./git init --author=Docli
{{< /highlight >}}

You've completed the tutorial! Thank you for reading and I hope you enjoyed it.

If you have any question, please don't hesitate to [open an issue](https://github.com/celicoo/docli/issues).
