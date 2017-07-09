package main

import (
	"fmt"
	"strconv"

	"./usr"
)

func listUserHomes() ([]string, error) {
	return nil, fmt.Errorf("listUserHomes() not yet implemented")
	lines, e := usr.LoadPasswdFrom(usr.PasswdPathLinux)
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
