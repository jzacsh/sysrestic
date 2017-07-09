package main

import (
	"log"
	"os"
)

const usageDoc string = `
  Name:
    sysrestic - an exclude-file joiner for system backups with restic

  Synopsis:
    sysreestic RESTIC_REPO EXCLUDE_FILE

  Description:
    Execs to restic[1] to backup / to RESTIC_REPO path with an automatically
    built list for restic's --exclude-file option.

  Outline:
    1. visits every $HOME on the system
    2. reads said $HOME's ~/.config/sysrestic.exclude or ~/.sysrestic.exclude
    3. creates a new exclude-file, unifying all $HOME's excludes w/EXCLUDE_FILE
    4. shells out to restic:
         restic backup \
            --repo RESTIC_REPO \
            --exclude-file /path/to/temporary/unified/exclude-list \
            /

  Reading Exclude Files:
    For both system and users' exclude files, empty files are okay.

    All lines in a user's exclude file are read as relative to their home.
    Leading slashes are ignored. Not much care has been taken beyond this to
    prevent bad things (eg: users may be able to exclude important files that do
    not belong to them using hard-link walks, like "../../../").

  [1]: https://restic.github.io
`

type resticCmd struct {
	ExcludeSysPath string
	ResticRepoPath string
}

func (c *resticCmd) parseExcludes() {
	panic("parseExcludes() not yet implemented: exclude.ParseHomeConf() in loop, then exclude.Build()")
}

func parseCli(args []string) error {
	panic("parseCli() not yet implemented")
}

func main() {
	log.Fatalf("not yet implemented\n")
	parseCli(os.Args[1:])
}
