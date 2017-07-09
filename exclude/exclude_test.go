package exclude

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func getFixtureDir() (string, error) {
	path := "../testdata"
	if f, err := os.Stat(path); (f != nil && !f.IsDir()) || err != nil {
		if f != nil && !f.IsDir() {
			return "", fmt.Errorf(
				"testdata was not a directory: %v", err)
		} else {
			return "", err
		}
	}
	return path, nil
}

func TestBuild(t *testing.T) {
	excludes := Build(
		[]string{"/proc/*", "/dev/*", "/tmp/*", "/lost+found"},
		[]string{"/home/alice/linus-tree", "/home/alice/build/", "/home/alice/mounts"},
	)
	expected := []string{
		"/proc/*", "/dev/*", "/tmp/*", "/lost+found",
		"/home/alice/linus-tree", "/home/alice/build/", "/home/alice/mounts",
	}
	for i, actual := range excludes {
		if actual == expected[i] {
			continue
		}

		t.Errorf(
			"expected line #%d to be '%s', but got '%s'",
			i+1, expected[i], actual)
	}
}

func TestParseHomeConf(t *testing.T) {
	fixtureDir, err := getFixtureDir()
	if err != nil {
		t.Fatalf("finding testdata: %v", err)
	}

	if _, err := ParseHomeConf("/dev/null"); err == nil {
		t.Errorf("seem OK with /dev/null")
	}

	aliceHome := filepath.Join(fixtureDir, "/home/alice")
	exc, err := ParseHomeConf(aliceHome)
	if err != nil {
		t.Errorf("unexpected problem with fixture, 'alice': %v", err)
	}

	// NOTE: keep this in sync with testdata under alice/ dir
	var expect []string
	for _, p := range []string{
		"/foo",
		"bar/*.txt",
		"baz/thing/",
		"/etc/",
	} {
		expect = append(expect, filepath.Join(aliceHome, p))
	}

	for i, path := range expect {
		if exc[i] == path {
			continue
		}
		t.Errorf(
			"expected line #%d to parse as:\n\t%s\n  but got:\n\t%s\n",
			i+1, path, exc[i])
	}
}
