package main

import (
	"errors"
	"time"
)

type Category struct {
	label string
	limit float64
	spent float64
}

type Budget struct {
	month      time.Month
	year       int32
	income     float64
	categories map[string]Category
}

func (b *Budget) AddCategory(c Category) error {
	_, ok := b.categories[c.label]
	if !ok {
		b.categories[c.label] = c
	} else {
		return errors.New("CategoryExistsError")
	}

	return nil
}

// use pointer to avoid copy? is that a thing? does it matter since im just collecting data?
func (b Budget) CalculateUnspent() float64 {
	total_spent := 0.00
	for _, v := range b.categories {
		total_spent += v.spent
		//fmt.Println(v.label, v.spent, "/", v.limit)
	}

	return b.income - total_spent
}

func (b *Budget) AddExpense(label string, amount float64) error {
	cat, ok := b.categories[label]
	if ok {
		cat.spent = cat.spent + amount
		b.categories[label] = cat
	} else {
		return errors.New("CategoryDoesntExistError")
	}

	return nil
}

func createBudget(month time.Month, year int32, income float64) Budget {
	return Budget{month, year, income, make(map[string]Category)}
}
