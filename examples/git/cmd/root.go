package cmd

import (
	"fmt"
	"log"

	"github.com/alecthomas/repr"
	"github.com/celicoo/docli"
	"github.com/celicoo/docli/pkg/args"
)

type Git struct {
	Clone   Clone
	Version Version
}

func (g *Git) Doc() string {
	return `usage: git <command>

commands:
  clone    clone a repository into a new directory
  version  print version
	
Use "git <command> help" for more information about the <command>.`
}

func (g *Git) Run() {
	repr.Println(g)
}

func (g *Git) Help() {
	fmt.Println(g.Doc())
}

func (g *Git) Error(err error) {
	switch err.(type) {
	case *args.InvalidArgumentError:
		// Ignore InvalidArgumentError.
		g.Run()
	default:
		log.Fatal(err)
	}
}

func Execute() {
	var g Git
	args := docli.Args()
	args.Bind(&g)
}
