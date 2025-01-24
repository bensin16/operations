package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type BudgetPageData struct {
	PageTitle   string
	BudgetSheet Budget
	GetDiff     func(float64, float64) float64
}

//func (b BudgetPageData) GetDiff(f1 float64, f2 float64) float64 {
//	fmt.Println("called GetDiff")
//	return f1 - f2
//}

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

	//	funcMap := template.FuncMap{
	//		"sub": func(f1 float64, f2 float64) float64 {
	//			return f1 - f2
	//		},
	//	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	getDiffFunc := func(f1 float64, f2 float64) float64 {
		return f1 - f2
	}
	tmpl.Execute(w, BudgetPageData{"Budget time", budget, getDiffFunc})
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", handle_index)

	http.ListenAndServe(":8090", nil)
}
