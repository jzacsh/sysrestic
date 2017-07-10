package exclude

import (
	"path/filepath"
	"testing"

	"../testdata"
)

const fixtureDir string = "../testdata"

func TestBuild(t *testing.T) {
	testdata.AssertFixtureDir(t, fixtureDir)
	sysExclude := filepath.Join(fixtureDir, "/etc/sysrestic.exclude")

	expected := []string{
		"/proc/*", "/dev/*", "/tmp/*", "/lost+found",
		"/home/alice/linus-tree", "/home/alice/build/", "/home/alice/mounts",
	}

	actual, e := Build(sysExclude, []string{
		"/home/alice/linus-tree",
		"/home/alice/build/",
		"/home/alice/mounts",
	})
	if e != nil {
		t.Fatalf("unexpected problem with Build(): %v", e)
	}
	if len(actual) != len(expected) {
		t.Fatalf("expected %d lines, but got %d", len(expected), len(actual))
	}

	for i, exp := range expected {
		if exp == actual[i] {
			continue
		}

		t.Errorf(
			"expected line #%d to be '%s', but got '%s'",
			i+1, exp, actual[i])
	}
}

func TestParseHomeConf(t *testing.T) {
	testdata.AssertFixtureDir(t, fixtureDir)

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

func TestParseHomeConf_AltPath(t *testing.T) {
	testdata.AssertFixtureDir(t, fixtureDir)

	janetHome := filepath.Join(fixtureDir, "/home/janet")
	exc, err := ParseHomeConf(janetHome)
	if err != nil {
		t.Errorf("unexpected problem with fixture, 'janet': %v", err)
	}

	// NOTE: keep this in sync with testdata under janet/ dir
	var expect []string
	for _, p := range []string{
		"/foo/bar/",
		".config/",
		"bioresearch.*.d",
	} {
		expect = append(expect, filepath.Join(janetHome, p))
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
