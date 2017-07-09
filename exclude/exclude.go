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
func Build(system string, homes ...[]string) ([]string, error) {
	var single []string

	f, e := os.Open(system)
	if e != nil {
		return nil, fmt.Errorf("opening %s: %v", system, e)
	}
	sc := bufio.NewScanner(f)
	var systemExcs []string
	for sc.Scan() {
		systemExcs = append(systemExcs, sc.Text())
	}
	if e := sc.Err(); e != nil {
		return nil, fmt.Errorf("reading %s: %v", system, e)
	}

	single = append(single, systemExcs...)
	for _, exclude := range homes {
		single = append(single, exclude...)
	}
	return single, nil
}

func findConf(homePath string) (*os.File, error) {
	const excludeName = "sysrestic.exclude"

	allowedConfs := []string{
		filepath.Join(homePath, ".config", excludeName),
		filepath.Join(homePath, fmt.Sprintf(".%s", excludeName)),
	}
	for _, path := range allowedConfs {
		conf, err := os.Open(path)
		if err != nil {
			continue
		}

		s, e := conf.Stat()
		if e != nil {
			return nil, fmt.Errorf("inspecting '%s': ", path, e)
		}

		if s.IsDir() {
			return nil, fmt.Errorf("'%s' should be a regular file, is a dir", path)
		}

		return conf, nil
	}

	return nil, fmt.Errorf(".config/ not top-level dot-file found")
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
	conf, err := findConf(homePath)
	if err != nil {
		return nil, err
	}
	var excludes []string

	scanner := bufio.NewScanner(conf)
	for scanner.Scan() {
		excludes = append(excludes, filepath.Join(homePath, scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return excludes, fmt.Errorf("reading %s: %s", conf.Name(), err)
	}
	return excludes, nil
}
