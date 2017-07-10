package file

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
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
