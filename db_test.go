package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBudgetCreateAndSave(t *testing.T) {
	dir, err := os.MkdirTemp("", "testdb")
	if err != nil {
		t.Fatalf("Couldn't create temp dir for testing")
	}
	defer os.RemoveAll(dir)

	dbFile := filepath.Join(dir, "test.db")

	var userId int64 = 1
	db := Database{Budgets: make(map[int64]Budget), FilePath: dbFile}

	b := createBudget(1, 2025, 5196.1)
	rentLabel := "Rent"
	b.AddCategory(Category{Label: rentLabel, Limit: 1000.00, Spent: 0.0})
	b.AddExpense(rentLabel, 1000.00)
	db.Budgets[userId] = b

	if err = db.Save(); err != nil {
		t.Fatalf("db failed to save")
	}

	actual, err := os.ReadFile(dbFile)

	// i dont really like that the category name is represented twice as the key and as a member of the struct. Maybe just drop the label from the struct?
	var expected string = "{\"month\":1,\"year\":2025,\"income\":5196.1,\"categories\":{\"Rent\":{\"label\":\"Rent\",\"limit\":1000,\"spent\":1000}}}"

	if err != nil {
		t.Fatalf("File couldn't be read")
	}

	if string(actual) != expected {
		t.Fatalf("Expected string did not match actual\nActual:   %s\nExpected: %s\n", actual, expected)
	}

}
