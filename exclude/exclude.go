package exclude

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// Joins all excludes into a single exclude list
//
// System excludes list is required, and arbitrary number of users' excludes are
// accepted.
func Build(system []string, homes ...[]string) []string {
	var single []string
	single = append(single, system...)
	for _, exclude := range homes {
		single = append(single, exclude...)
	}
	return single
}

// Read exclude paths for a given user's home directory
//
// homePath: is a string indicatging path to single user's home "/home/bob" or
// "/usrhomes/disk2/alic/". The exclude file found under homePath will be read
// as follows:
// - ascii encoded text
// - each line indicates a path to be excluded
// - all lines are read as relative to $HOME; leading slashes ignored
//
// Return is a list of absolute paths that should be excluded
func ParseHomeConf(homePath string) ([]string, error) {
	confPath := filepath.Join(homePath, ".config/sysrestic.exclude")
	conf, err := os.Open(confPath)
	if err != nil {
		return nil, err
	}
	var excludes []string

	scanner := bufio.NewScanner(conf)
	for scanner.Scan() {
		excludes = append(excludes, filepath.Join(homePath, scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return excludes, fmt.Errorf("reading %s: %s", confPath, err)
	}
	return excludes, nil
}
