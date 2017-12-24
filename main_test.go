package main

import "testing"

func TestCheck(t *testing.T) {
	err := check("./testdata/pkgs/badpkg/...", []string{"go/ast", "go/token"})
	if err == nil {
		t.Fatal("expected an error here")
	}
	expect := "./testdata/pkgs/badpkg is using banned dependencies go/ast, go/token"
	got := err.Error()
	if expect != got {
		t.Fatalf("expected %s, got %s", expect, got)
	}
}

func TestCheckGoodCode(t *testing.T) {
	// check package
	err := check("./testdata/pkgs/goodpkg/...", []string{"go/ast"})
	if err != nil {
		t.Fatal("got an unexpected error:", err)
	}
	// check specific folder
	err = check("./testdata/pkgs/goodpkg", []string{"go/ast"})
	if err != nil {
		t.Fatal("got an unexpected error:", err)
	}
}

func TestCheckPkgBadFile(t *testing.T) {
	err := checkPkg("./testdata/pkgs/DNE", []string{"fmt"})
	if err == nil {
		t.Fatal("expected an error here")
	}
	expect := "failed to parse pkg: ./testdata/pkgs/DNE: open ./testdata/pkgs/DNE: no such file or directory"
	got := err.Error()
	if expect != got {
		t.Fatalf("expected %s, got %s", expect, got)
	}
}
