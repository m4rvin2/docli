package cmd

import (
	"fmt"
	"log"
)

const version = "0.1.0"

type Version struct{}

func (v *Version) Doc() string {
	return ""
}

func (v *Version) Run() {
	fmt.Printf("git version %s\n", version)
}

func (v *Version) Help() {
	v.Run()
}

func (v *Version) Error(err error) {
	log.Fatal(err)
}
