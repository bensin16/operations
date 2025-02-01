package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type BudgetPageData struct {
	PageTitle   string
	BudgetSheet Budget
}

type ExpenseForm struct {
	Category string
	Planned  float64
	Actual   float64
}

func GetDiff(f1 float64, f2 float64) float64 {
	return f1 - f2
}

func add_basic_categories(budget *Budget) {
	err := budget.AddCategory(Category{"Rent", 1375.00, 0})
	if err != nil {
		fmt.Println("Category creation failed")
		return
	}

	err = budget.AddCategory(Category{"Grocery", 500.00, 0})
	if err != nil {
		fmt.Println("Category creation failed")
		return
	}

	err = budget.AddCategory(Category{"Electric", 60.00, 0})
	if err != nil {
		fmt.Println("Category creation failed")
		return
	}
}

func handle_index(w http.ResponseWriter, req *http.Request) {
	budget := createBudget(1, 2025, 5195.10)
	add_basic_categories(&budget)

	err := budget.AddExpense("Rent", 1375.00)
	if err != nil {
		fmt.Println("Expense creation failed")
		return
	}

	err = budget.AddExpense("Grocery", 63.32)
	if err != nil {
		fmt.Println("Grocery creation failed")
		return
	}

	funcMap := template.FuncMap{
		"diff": GetDiff,
	}

	tmpl, err := template.New("index.html").Funcs(funcMap).ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, BudgetPageData{"Budget time", budget})
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", handle_index)

	log.Fatal(http.ListenAndServe(":8090", nil))
}
