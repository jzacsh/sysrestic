package file

import (
	"fmt"
	"os"
)

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

// IsReadableDir determines if path seems to point to a readable directory.
// Reminiscent of a shell `[ -e path && -d path ]` test builtin.
//
// False returns include an error for justification.
func IsReadableDir(path string) (bool, error) {
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

// IsReadableFile determines if path seems to point to a readable file.
// Reminiscent of a shell `[ -e path ]` test builtin.
//
// False returns include an error for justification.
func IsReadableFile(path string) (bool, error) {
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
