package main

import (
	"os"
)

// its just a single budget for now but I want it to be a list of them
type Database struct {
	Budgets  map[int64]Budget
	FilePath string
}

func (d *Database) Save() error {
	d1 := []byte("hello\nworld\n")
	err := os.WriteFile(d.FilePath, d1, 0644)
	if err != nil {
		return err
	}

	return nil
}
