package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"./exclude"
	"./file"
)

type resticCmd struct {
	ExcludeSysPath  string
	ResticRepoPath  string
	BackupTarget    string
	Err             *log.Logger
	UnifiedExcludes string
	Excludes        []string
}

func (c *resticCmd) String() string {
	return fmt.Sprintf(
		"RESTIC_REPO:  '%s'\nEXCLUDE_FILE:  '%s'\n",
		c.ResticRepoPath, c.ExcludeSysPath)
}

func (c *resticCmd) parseExcludes() error {
	homes, e := listHumanUserHomes_Linux()
	if e != nil {
		return fmt.Errorf("finding $HOMEs: %v", e)
	}

	var excs [][]string
	for _, home := range homes {
		excludes, e := exclude.ParseHomeConf(home)
		if e != nil {
			return fmt.Errorf("parsing %s excludes: %v", home, e)
		}
		excs = append(excs, excludes)
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
	return file.WriteAsciiLines(c.Excludes, c.UnifiedExcludes)
}

func (c *resticCmd) runBackup() error {
	cmd := exec.Command(
		"restic", "backup",
		"--repo", c.ResticRepoPath,
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
	r.Err = lg

	if e := r.parseExcludes(); e != nil {
		lg.Fatalf("excludes: %v\n", e)
	}

	fmt.Printf("%d excludes written to %s\n", len(r.Excludes), r.UnifiedExcludes)

	if e := r.runBackup(); e != nil {
		r.Err.Fatalf("restic: %v\n", e)
	}
}
