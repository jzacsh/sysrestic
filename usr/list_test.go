package usr

import (
	"os/user"
	"path/filepath"
	"testing"

	"../test"
)

const fixtureDir string = "../testdata"

func TestLoadPasswdFrom(t *testing.T) {
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
	actual, e := LoadPasswdFrom(passwd)
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

func TestLoadPasswdFrom_EmptyFile(t *testing.T) {
	actual, e := LoadPasswdFrom("/dev/null")
	if e != nil {
		t.Fatalf("unexpected, loading test empty passwd: %s", e)
	}

	if len(actual) != 0 {
		t.Fatalf("expected empty result, got %d lines", len(actual))
	}
}

func TestListUsers(t *testing.T) {
	test.AssertFixtureDir(t, fixtureDir)

	passwd, e := LoadPasswdFrom(filepath.Join(fixtureDir, "/etc/passwd"))
	if e != nil {
		t.Fatalf("unexpected, loading test passwd file: %s", e)
	}

	// NOTE: keep synced with testdata/etc/passwd
	expect := []user.User{
		user.User{Uid: "0", Gid: "0", Username: "root", Name: "root", HomeDir: "/root"},
		user.User{Uid: "1000", Gid: "1000", Username: "alice", Name: "Alice", HomeDir: "/home/alice"},
		user.User{Uid: "1001", Gid: "1001", Username: "janet", Name: "Janet", HomeDir: "/home/janet"},
		user.User{Uid: "65534", Gid: "65534", Username: "nobody", Name: "nobody", HomeDir: "/nonexistent"},
		user.User{Uid: "117", Gid: "124", Username: "pulse", Name: "PulseAudio daemon", HomeDir: "/var/run/pulse"},
	}

	actual, e := ListUsers(passwd)
	if e != nil {
		t.Fatalf("unexpected, parsing passwd lines: %s", e)
	}

	ueq := func(a, b user.User) bool {
		return a.Uid == b.Uid &&
			a.Gid == b.Gid &&
			a.Username == b.Username &&
			a.Name == b.Name &&
			a.HomeDir == b.HomeDir
	}

	if len(actual) != len(expect) {
		t.Fatalf("expected %d users, but got %d", len(expect), len(actual))
	}

	for i, u := range expect {
		if ueq(u, actual[i]) {
			continue
		}
		t.Errorf("expected user:\n\t%v\nbut actually got user:\n\t%v\n", u, actual[i])
	}
}
