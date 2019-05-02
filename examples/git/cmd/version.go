package cmd

import (
	"fmt"
	"log"
)

const version = "0.1.0"

type Version struct{}

func (c *Version) Doc() string {
	return ""
}

func (c *Version) Run() {
	fmt.Printf("git version %s\n", version)
}

func (v *Version) Help() {
	v.Run()
}

func (c *Version) Error(err error) {
	log.Fatal(err)
}
