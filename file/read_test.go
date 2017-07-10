package file

import (
	"path/filepath"
	"testing"

	"../test"
)

const fixtureDir string = "../testdata"

func TestReadLines(t *testing.T) {
	test.AssertFixtureDir(t, fixtureDir)

	// NOTE: keep synced with testdata/etc/passwd
	expect := []string{
		"root:x:0:0:root:/root:/bin/bash",
		"alice:x:1000:1000:Alice,,,:/home/alice:/bin/bash",
		"janet:x:1001:1001:Janet,,,:/home/janet:/bin/bash",
		"nobody:x:65534:65534:nobody:/nonexistent:/usr/sbin/nologin",
		"pulse:x:117:124:PulseAudio daemon,,,:/var/run/pulse:/bin/false",
	}

	passwd := filepath.Join(fixtureDir, "/etc/passwd")
	actual, e := ReadLines(passwd)
	if e != nil {
		t.Fatalf("unexpected, loading test passwd file: %s", e)
	}

	if len(actual) != len(expect) {
		t.Fatalf("expected %d lines, but got %d", len(expect), len(actual))
	}

	for i, line := range expect {
		if actual[i] == line {
			continue
		}

		t.Errorf("line #%d expected '%s', got '%s'", i, line, actual[i])
	}
}

func TestReadLines_EmptyFile(t *testing.T) {
	actual, e := ReadLines("/dev/null")
	if e != nil {
		t.Fatalf("unexpected, loading test empty passwd: %s", e)
	}

	if len(actual) != 0 {
		t.Fatalf("expected empty result, got %d lines", len(actual))
	}
}
