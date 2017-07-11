package file

import (
	"path/filepath"
	"testing"

	"github.com/jzacsh/sysrestic/testdata"
)

func TestWriteASCIILines(t *testing.T) {
	testdata.AssertFixtureDir(t, fixtureDir)

	p := filepath.Join(fixtureDir, "/non/existent/dir/not/possible/to/write/to/")
	if WriteASCIILines([]string{"foo"}, p) == nil {
		t.Errorf("expected failure writing to: %s", p)
	}

	p = filepath.Join(fixtureDir, "/output/writetest.txt")
	if e := WriteASCIILines([]string{"foo", "bar", "baz"}, p); e != nil {
		t.Errorf("unexpected failure writing to %s: %v", p, e)
	}

	// TODO diff against p against ex:
	// ex = filepath.Join(fixtureDir, "/home/janet/pkgout/write.WriteASCIILines_foobarbaz")
}
