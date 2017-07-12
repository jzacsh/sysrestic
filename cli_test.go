package main

import (
	"path/filepath"
	"testing"

	"github.com/jzacsh/sysrestic/testdata"
)

const fixtureDir string = "testdata"

func TestParseCliHelp(t *testing.T) {
	helps := []string{
		"h", "-h", "--h",
		"help", "-help", "--help",
	}

	for _, help := range helps {
		if cmd, err := parseCli([]string{help}); cmd != nil || err != nil {
			t.Errorf("expected `%s` flag to trigger help doc behavior", help)
		}
	}
}

func TestParseCliMissingArgs(t *testing.T) {
	testdata.AssertFixtureDir(t, fixtureDir)

	clis := [][]string{
		[]string{},
		[]string{
			filepath.Join(fixtureDir, "/etc/"),
			filepath.Join(fixtureDir, "/etc/sysrestic.exclude"),
		},
		[]string{"/some/repo", "/some/exclude", "--help"},
	}

	for _, args := range clis {
		cmd, err := parseCli(args)
		if cmd == nil && err != nil {
			t.Logf("properly reported problem; got: %v\n", err)
			continue
		}
		t.Errorf(
			"expected cli `%v` to be caught; got no error & cmd:\n%v\n",
			args, cmd)
	}
}

func TestParseCliBadArgs(t *testing.T) {
	testdata.AssertFixtureDir(t, fixtureDir)

	clis := [][]string{
		[]string{
			"file types inverted: repo first, then file",
			filepath.Join(fixtureDir, "/etc/passwd"),
			filepath.Join(fixtureDir, "/etc"),
		},
		[]string{
			"repo ok, but file non-existent",
			filepath.Join(fixtureDir, "/etc/"),
			filepath.Join(fixtureDir, "/some/exclude"),
		},
	}

	for _, args := range clis {
		if cmd, err := parseCli(args[1:]); cmd != nil || err == nil {
			t.Errorf(
				"expected cli `%v` to fail because %s, but got cmd:\n%v\n",
				args[1:], args[0], cmd)
		}
	}
}

func TestParseCliOK(t *testing.T) {
	testdata.AssertFixtureDir(t, fixtureDir)

	args := []string{filepath.Join(fixtureDir, "/etc/sysrestic.exclude")}

	cmd, err := parseCli(args)

	if err != nil {
		t.Errorf(
			"unexpected failure of cli `%v`; got: %v\n",
			args, err)
	}

	if cmd == nil {
		t.Errorf("empty CMD for valid commandline")
	}
	if cmd.ExcludeSysPath != args[0] {
		t.Errorf("bad CMD; expected EXCLUDE_FILE to be %s, but got '%s'",
			args[0], cmd.ExcludeSysPath)
	}
}
