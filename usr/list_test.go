package usr

import (
	"testing"

	"../test"
)

const fixtureDir string = "../testdata"

func TestLoadPasswdFrom(t *testing.T) {
	test.AssertFixtureDir(t, fixtureDir)
	t.Fatalf("LoadPasswdFrom() not yet tested")
}

func TestListUsers(t *testing.T) {
	test.AssertFixtureDir(t, fixtureDir)
	t.Fatalf("ListUsers() not yet tested")
}
