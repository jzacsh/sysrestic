package usr

import (
	"fmt"
	"os/user"
	"strings"
)

// Location on GNU/Linux systems where user-data is stored
const PasswdPathLinux string = "/etc/passwd"

type UsrList []user.User

// Parses a /etc/passwd formatted line, and produces os/user.User object
// eg: given "alice:x:1000:1000:Alice,,,:/home/alice:/bin/bash", produces:
//   user.User{
//     Uid: "1000",
//     Gid: "1000",
//     Username: "alice",
//     Name: "Alice",
//     HomeDir: "/home/alice",
//   }
//
// If an error is returned, it is due to format issues
func parseUserLine(line string) (*user.User, error) {
	// NOTE: inspired mostly by go std lib's os/user pkg internal Lookup() logic
	//       Unfortunately none of is written with API surface area

	parts := strings.SplitN(line, ":", 7)
	if len(parts) < 6 {
		// TODO(zacsh) revisit; does this actually happen?
		return nil, fmt.Errorf("unexpectedly short line; got %d parts", len(parts))
	}

	u := &user.User{
		Username: parts[0],
		Uid:      parts[2],
		Gid:      parts[3],
		Name:     parts[4],
		HomeDir:  parts[5],
	}

	// Only want the first field
	if i := strings.Index(u.Name, ","); i >= 0 {
		u.Name = u.Name[:i]
	}

	return u, nil
}

// List users according to a system-listing like that of /etc/passwd on
// gnu/linux
//
// lines is the contents of /etc/passwd (or some similar output) where each
// element was originally distinguished by a line-break.
func ListUsers(lines []string) (UsrList, error) {
	var list UsrList

	for i, line := range lines {
		u, e := parseUserLine(line)
		if e != nil {
			return nil, fmt.Errorf("parsing line #%d: %s", i+1, e)
		}

		list = append(list, *u)
	}

	return list, nil
}
