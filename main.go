package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/jzacsh/sysrestic/exclude"
	"github.com/jzacsh/sysrestic/file"
)

type resticCmd struct {
	ExcludeSysPath  string
	BackupTarget    string
	Err             *log.Logger
	UnifiedExcludes string
	Excludes        []string
	Users           int
	UserExcludes    int
}

func (c *resticCmd) String() string {
	return fmt.Sprintf("EXCLUDE_FILE:  '%s'\n", c.ExcludeSysPath)
}

func (c *resticCmd) parseExcludes() error {
	homes, e := listHumanUserHomesLinux()
	if e != nil {
		return fmt.Errorf("finding $HOMEs: %v", e)
	}
	c.Users = len(homes)

	resticRepo := os.Getenv("RESTIC_REPOSITORY")
	if len(resticRepo) == 0 {
		return fmt.Errorf("$RESTIC_REPOSITORY empty")
	}

	// ensure restic excludes its own repo
	excs := [][]string{[]string{resticRepo}}

	for _, home := range homes {
		excludes, e := exclude.ParseHomeConf(home)
		if e != nil {
			return fmt.Errorf("parsing %s excludes: %v", home, e)
		}
		excs = append(excs, excludes)
		c.UserExcludes++
	}

	unified, e := exclude.Build(c.ExcludeSysPath, excs...)
	if e != nil {
		return e
	}
	c.Excludes = unified

	if len(c.Excludes) == 0 {
		return nil
	}

	f, e := ioutil.TempFile("" /*default*/, "sysrestic-unified-excludes_")
	if e != nil {
		return fmt.Errorf("failed to start tempfile for excludes: %v", e)
	}
	defer f.Close()
	c.UnifiedExcludes = f.Name()

	// TODO(zacsh) remove this block and convert all `string` signatures for
	// exclude-file handling to `byte`, since we just want ascii
	return file.WriteASCIILines(c.Excludes, c.UnifiedExcludes)
}

func printResticVersion() {
	cmd := exec.Command("restic", "version")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if e := cmd.Run(); e != nil {
		panic(fmt.Errorf("checking restic's version: %v", e))
	}
}

func (c *resticCmd) runBackup() error {
	cmd := exec.Command(
		"restic", "backup",
		"--one-file-system",
		"--exclude-file", c.UnifiedExcludes,
		c.BackupTarget)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	lg := log.New(os.Stderr, "sysrestic: ", 0)
	r, e := parseCli(os.Args[1:])
	if e != nil {
		lg.Fatalf("parsing command: %v\n", e)
	}
	if r == nil {
		fmt.Printf(usageDoc)
		os.Exit(0)
	}
	r.Err = lg

	if e := r.parseExcludes(); e != nil {
		lg.Fatalf("excludes: %v\n", e)
	}

	fmt.Printf(
		"%d excludes from %d of %d users written to %s\n",
		len(r.Excludes), r.UserExcludes, r.Users, r.UnifiedExcludes)

	printResticVersion()

	if e := r.runBackup(); e != nil {
		r.Err.Fatalf("restic: %v\n", e)
	}

	fmt.Printf("Restic exited OK. Cleaning up...")
	if e := os.Remove(r.UnifiedExcludes); e != nil {
		fmt.Printf("\n")
		r.Err.Fatalf("tmpfile removal: %v\n", e)
	}
	fmt.Printf(" done.\n")
}
