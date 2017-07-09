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
	log.Fatalf("parseExcludes() not yet implemented: exclude.ParseHomeConf() in loop, then exclude.Build()")
}

func main() {
	lg := log.New(os.Stderr, "sysrestic: ", 0)
	r, e := parseCli(os.Args[1:])
	if e != nil {
		lg.Fatalf("parsing command: %s\n", e)
	}
	lg.Fatalf("not yet implemented, but got: %s\n", r)
}
