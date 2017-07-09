package main

import (
	"fmt"
	"log"
	"os"
)

type resticCmd struct {
	ExcludeSysPath string
	ResticRepoPath string
	Log            *log.Logger
}

func (c *resticCmd) String() string {
	return fmt.Sprintf(
		"RESTIC_REPO:  '%s'\nEXCLUDE_FILE:  '%s'\n",
		c.ResticRepoPath, c.ExcludeSysPath)
}

func (c *resticCmd) parseExcludes() {
	c.Log.Fatalf("parseExcludes() not yet implemented: exclude.ParseHomeConf() in loop, then exclude.Build()")
}

func main() {
	lg := log.New(os.Stderr, "sysrestic: ", 0)
	r, e := parseCli(os.Args[1:])
	if e != nil {
		lg.Fatalf("parsing command: %s\n", e)
	}
	r.Log = lg

	r.Log.Fatalf("not yet implemented, but got:\n%s\n", r)
}
