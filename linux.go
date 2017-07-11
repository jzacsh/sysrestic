package main

import (
	"fmt"
	"strconv"

	"github.com/jzacsh/sysrestic/file"
	"github.com/jzacsh/sysrestic/usr"
)

// TODO(zacsh) figure out how/if to test
func listHumanUserHomes_Linux() ([]string, error) {
	lines, e := file.ReadLines(usr.PasswdPathLinux)
	if e != nil {
		return nil, fmt.Errorf("loading %s: %s", usr.PasswdPathLinux, e)
	}

	usrs, e := usr.ListUsers(lines)
	if e != nil {
		return nil, fmt.Errorf("malformed %s: %s", usr.PasswdPathLinux, e)
	}

	var homes []string
	for _, u := range usrs {
		if u.Username == "nobody" || len(u.HomeDir) == 0 {
			continue
		}

		uid, e := strconv.Atoi(u.Uid)
		if e != nil {
			return nil, fmt.Errorf(
				"malformed user='%s' in %s: %s", u.Username, usr.PasswdPathLinux, e)
		}
		if uid < 1000 {
			continue
		}

		homes = append(homes, u.HomeDir)
	}
	return homes, nil
}
