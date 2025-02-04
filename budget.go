package main

import (
	"errors"
	"time"
)

type Category struct {
	Label string  `json:"label"`
	Limit float64 `json:"limit"`
	Spent float64 `json:"spent"`
}

type Budget struct {
	Month      time.Month          `json:"month"`
	Year       int32               `json:"year"`
	Income     float64             `json:"income"`
	Categories map[string]Category `json:"categories"`
}

func (b *Budget) AddCategory(c Category) error {
	_, ok := b.Categories[c.Label]
	if !ok {
		b.Categories[c.Label] = c
	} else {
		return errors.New("CategoryExistsError")
	}

	return nil
}

// use pointer to avoid copy? is that a thing? does it matter since im just collecting data?
func (b Budget) CalculateUnspent() float64 {
	total_spent := 0.00
	for _, v := range b.Categories {
		total_spent += v.Spent
		//fmt.Println(v.Label, v.Spent, "/", v.Limit)
	}

	return b.Income - total_spent
}

func (b *Budget) AddExpense(label string, amount float64) error {
	cat, ok := b.Categories[label]
	if ok {
		cat.Spent = cat.Spent + amount
		b.Categories[label] = cat
	} else {
		return errors.New("CategoryDoesntExistError")
	}

	return nil
}

func createBudget(month time.Month, year int32, income float64) Budget {
	return Budget{month, year, income, make(map[string]Category)}
}
