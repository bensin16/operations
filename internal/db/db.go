package db

import (
	"encoding/json"
	"os"

	"doobir.net/operations/internal/budget"
)

// its just a single budget for now but I want it to be a list of them
type Database struct {
	Budgets  map[int64]budget.Budget
	FilePath string
}

// whats my writing format? Let's try a json encoder
func (d *Database) Save() error {

	encodedBudget, err := json.Marshal(d.Budgets[1])
	if err != nil {
		return err
	}

	d1 := encodedBudget

	f, err := os.Create(d.FilePath)
	if err != nil {
		panic(err) // find better way to log error then return value
	}
	defer f.Close()

	_, err = f.Write(d1) // _ = bytes written
	if err != nil {
		return err
	}

	return nil
}
