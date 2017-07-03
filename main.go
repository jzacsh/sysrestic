package main

import (
	"log"
	"os"
)

type resticCmd struct {
	ExcludeSysPath string
	ResticRepoPath string
}

func (c *resticCmd) parseExcludes() {
	panic("parseExcludes() not yet implemented: exclude.ParseHomeConfs() in loop, then exclude.Build()")
}

func parseCli(args []string) error {
	panic("parseCli() not yet implemented")
}

func main() {
	log.Fatalf("not yet implemented\n")
	parseCli(os.Args[1:])
}
