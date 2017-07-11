package file

import (
	"path/filepath"
	"testing"

	"github.com/jzacsh/sysrestic/testdata"
)

func TestIsReadableFile(t *testing.T) {
	testdata.AssertFixtureDir(t, fixtureDir)

	p := filepath.Join(fixtureDir, "/etc/")
	if ok, _ := IsReadableFile(p); ok {
		t.Errorf("false positive: %s is directory", p)
	}

	p = filepath.Join(fixtureDir, "/etc/passwd")
	if ok, e := IsReadableFile(p); !ok {
		t.Errorf("%s should be readable file; got: %v", p, e)
	}
}

func TestIsReadableDir(t *testing.T) {
	testdata.AssertFixtureDir(t, fixtureDir)

	p := filepath.Join(fixtureDir, "/etc/passwd")
	if ok, _ := IsReadableDir(p); ok {
		t.Errorf("false positive: %s is file", p)
	}

	p = filepath.Join(fixtureDir, "/etc/")
	if ok, e := IsReadableDir(p); !ok {
		t.Errorf("%s should be readable dir; got: %v", p, e)
	}
}
