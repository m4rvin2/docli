package cmd

import (
	"fmt"
	"log"

	"github.com/celicoo/docli/v3"
	"github.com/celicoo/docli/v3/pkg/arguments"
)

const version = "0.0.1"

type Git struct {
	Clone   Clone
	Version bool
}

func (g *Git) Doc() string {
	return `usage: git <command>

commands:
  c, clone       clone a repository into a new directory
 
arguments:

  -v, --version  print version
	
Use "git <command> help" for more information about the <command>.`
}

func (g *Git) Run() {
	if g.Version {
		fmt.Println(version)
		return
	}
	fmt.Println("Hello, world!")
}

func (g *Git) Help() {
	fmt.Println(g.Doc())
}

func (g *Git) Error(err error) {
	switch err.(type) {
	case *arguments.InvalidArgumentError:
		// Ignore InvalidArgumentError.
		g.Run()
	default:
		log.Fatal(err)
	}
}

func Execute() {
	var g Git
	docli.Arguments.Bind(&g)
}
