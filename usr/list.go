package usr

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
)

// Location on GNU/Linux systems where user-data is stored
const PasswdPathLinux string = "/etc/passwd"

type UsrList []user.User

// Produces output intended as a `ListUsers` parameter.
//
// path should be a filepath, eg: `PasswdPathLinux`
//
// contents of path's file are taken to be newline-delimeted, and loaded as
// elements of the returned slice.
func LoadPasswdFrom(path string) ([]string, error) {
	f, e := os.Open(path)
	if e != nil {
		return nil, fmt.Errorf("opening: %s", e)
	}

	s := bufio.NewScanner(f)

	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	if e := s.Err(); e != nil {
		return nil, fmt.Errorf("reading: %s", e)
	}

	return lines, nil
}

// List users according to a system-listing like that of /etc/passwd on
// gnu/linux
//
// lines is the contents of /etc/passwd (or some similar output) where each
// element was originally distinguished by a line-break.
func ListUsers(lines []string) (UsrList, error) {
	return nil, fmt.Errorf("ListUsers(...) not yet implemented; but got: %s", lines)
}
