package main

import (
	"fmt"
	"log"
	"os"
)

type resticCmd struct {
	ExcludeSysPath string
	ResticRepoPath string
	Err            *log.Logger
}

func (c *resticCmd) String() string {
	return fmt.Sprintf(
		"RESTIC_REPO:  '%s'\nEXCLUDE_FILE:  '%s'\n",
		c.ResticRepoPath, c.ExcludeSysPath)
}

func (c *resticCmd) parseExcludes() {
	c.Err.Fatalf("parseExcludes() not yet implemented: exclude.ParseHomeConf() in loop, then exclude.Build()")
}

func main() {
	lg := log.New(os.Stderr, "sysrestic: ", 0)
	r, e := parseCli(os.Args[1:])
	if e != nil {
		lg.Fatalf("parsing command: %s\n", e)
	}
	r.Err = lg

	r.Err.Fatalf("not yet implemented, but got:\n%s\n", r)
}
