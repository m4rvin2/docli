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
	// Just print and exit, regardless.
	log.Fatal(err)
}
