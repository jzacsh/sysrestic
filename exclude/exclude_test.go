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
	t.Errorf("tests of Build() not yet implemented")
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

	t.Logf("%d exclude paths\n", len(exc)) // TODO finish test
	t.Errorf("tests of ParseHomeConf() not yet implemented")
}
