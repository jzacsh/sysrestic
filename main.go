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
	// TODO: call listUserHomes()
}

func main() {
	lg := log.New(os.Stderr, "sysrestic: ", 0)
	r, e := parseCli(os.Args[1:])
	if e != nil {
		lg.Fatalf("parsing command: %s\n", e)
	}
	r.Err = lg

	// TODO(zacsh) remaining steps to implement:
	// 1. gets listing of every $HOME on the system
	//    a. getent passwd
	//       to get real/human users on the machine, is a pain in the ass; here's
	//       what it is in shell:
	//
	//           while read usr uid hm; do
	//             [[ "$uid" -ge 1000 && "$usr" != nobody ]] || continue
	//             printf '%s[%d]: %s\n' "$usr" $uid "$hm"
	//           done < <(getent passwd | awk -F : '{print $1 "\t" $3 "\t" $6 }' )
	//
	//       TODO(zacsh) GOOS: find OSX-way to do this & add ifdef
	//    b. discard lines without a home
	// 2. reads said $HOME's ~/.config/sysrestic.exclude or ~/.sysrestic.exclude
	//    ie: call exclude.ParseHomeConf()
	// 3. creates a new exclude-file, unifying all $HOME's excludes w/EXCLUDE_FILE
	//    ie: call exclude.Build(c.ExcludeSysPath, ....)
	// 4. shells out to restic:
	//    a. open tempfile
	//    b. dump in exclude.Build(...)'s result
	//    c. os.Exec(...) .Run() with tempfile as arg

	r.Err.Fatalf("not yet implemented, but got:\n%s\n", r)
}
