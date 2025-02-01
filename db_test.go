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

	var userId int64 = 1
	db := Database{Budgets: make(map[int64]Budget), FilePath: dbFile}
	db.Budgets[userId] = createBudget(1, 2025, 5196.1)

	if err = db.Save(); err != nil {
		t.Fatalf("db failed to save")
	}

	actual, err := os.ReadFile(dbFile)

	var expected string = "{\"Month\":1,\"Year\":2025,\"Income\":5196.1,\"Categories\":{}}"

	if err != nil {
		t.Fatalf("File couldn't be read")
	}

	if string(actual) != expected {
		// look at that one page for how to print out the values
		t.Fatalf("Expected string did not match actual")
	}

}
