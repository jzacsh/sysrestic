package test

import (
	"os"
	"testing"
)

func AssertFixtureDir(t *testing.T, path string) {
	if f, err := os.Stat(path); (f != nil && !f.IsDir()) || err != nil {
		if f != nil && !f.IsDir() {
			t.Fatalf("fixture path '%s' not a directory: %v", err)
		}
	}
}
