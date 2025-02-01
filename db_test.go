package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSave(t *testing.T) {
	dir, err := os.MkdirTemp("", "testdb")
	if err != nil {
		t.Fatalf("Couldn't create temp dir for testing")
	}
	defer os.RemoveAll(dir)

	dbFile := filepath.Join(dir, "test.db")

	db := Database{FilePath: dbFile}

	if err = db.Save(); err != nil {
		t.Fatalf("db failed to save")
	}

	// i guess just check that i can read the file for now, can make better test for actually checking for file contents/"integration" tests for loading a budget when i get there
	_, err = os.ReadFile(dbFile)

	if err != nil {
		t.Fatalf("File couldn't be read")
	}
}
