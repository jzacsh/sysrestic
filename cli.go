package main

import (
	"fmt"
	"os"
	"strings"
)

const defaultBackupTarget string = "/"

const usageDoc string = `sysrestic - an exclude-file joiner for system backups with restic

Synopsis:
  sysreestic [help] RESTIC_REPO EXCLUDE_FILE

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

func looksLikeHelp(arg string) bool {
	return arg == "h" || arg == "-h" || arg == "--h" ||
		arg == "help" || arg == "-help" || arg == "--help"
}

func getReadableStat(path string) (os.FileInfo, error) {
	f, e := os.Open(path)
	if e != nil {
		return nil, e
	}
	defer f.Close()
	s, e := f.Stat()
	if e != nil {
		return nil, e
	}
	return s, nil
}

func isReadableDir(path string) (bool, error) {
	s, e := getReadableStat(path)
	if e != nil {
		return false, e
	}

	isDir := s.IsDir()
	if isDir {
		return true, nil
	}
	return false, fmt.Errorf("%s not a directory", path)
}

func isReadableFile(path string) (bool, error) {
	s, e := getReadableStat(path)
	if e != nil {
		return false, e
	}

	isFile := !s.IsDir()
	if isFile {
		return true, nil
	}
	return false, fmt.Errorf("%s is a directory", path)
}

func parseCli(args []string) (*resticCmd, error) {
	if len(args) != 2 {
		if len(args) == 1 && looksLikeHelp(args[0]) {
			fmt.Printf(usageDoc)
			os.Exit(0)
		}
		return nil, fmt.Errorf("must provide 2 args, got %d", len(args))
	}

	r := &resticCmd{
		ResticRepoPath: strings.TrimSpace(args[0]),
		ExcludeSysPath: strings.TrimSpace(args[1]),
		BackupTarget:   defaultBackupTarget,
	}
	if is, e := isReadableDir(r.ResticRepoPath); !is {
		return nil, fmt.Errorf("RESTIC_REPO not a readable dir: %s", e)
	}
	if is, e := isReadableFile(r.ExcludeSysPath); !is {
		return nil, fmt.Errorf("EXCLUDE_FILE not a readable file: %s", e)
	}

	return r, nil
}
