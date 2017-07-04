package exclude

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func Build(system []string, homes ...[]string) ([]string, error) {
	panic("Build() not yet implemented")
}

// Read exclude paths for a given user's home directory
//
// homePath: is a string indicatging path to single user's home "/home/bob" or
// "/usrhomes/disk2/alic/" is as follows:
// - ascii encoded text
// - relative paths to $HOME; leading slashes will be ignored
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
