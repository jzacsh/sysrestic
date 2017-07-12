package main

import (
	"fmt"
	"strings"

	"github.com/jzacsh/sysrestic/file"
)

const defaultBackupTarget string = "/"

const usageDoc string = `sysrestic - an exclude-file joiner for system backups with restic

Synopsis:
  sysreestic [help] EXCLUDE_FILE

Description:
  Execs to restic[1] to backup / to $RESTIC_REPOSITORY path with an
  automatically built list for restic's --exclude-file option.

Outline:
  1. visits every $HOME on the system
  2. reads said $HOME's ~/.config/sysrestic.exclude or ~/.sysrestic.exclude
  3. creates a new exclude-file, unifying all $HOME's excludes w/EXCLUDE_FILE
  4. ensures $RESTIC_REPOSITORY is set
  5. shells out to restic:
       restic backup --exclude-file /path/to/temporary/unified/exclude-list /

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

func parseCli(args []string) (*resticCmd, error) {
	if len(args) != 1 || looksLikeHelp(args[0]) {
		if len(args) == 1 && looksLikeHelp(args[0]) {
			return nil, nil
		}
		return nil, fmt.Errorf("must provide EXCLUDE_FILE, got %d args", len(args))
	}

	r := &resticCmd{
		BackupTarget:   defaultBackupTarget,
		ExcludeSysPath: strings.TrimSpace(args[0]),
	}
	if is, e := file.IsReadableFile(r.ExcludeSysPath); !is {
		return nil, fmt.Errorf("EXCLUDE_FILE, '%s' not a readable file: %s", r.ExcludeSysPath, e)
	}

	return r, nil
}
