package main

import (
	"html/template"
	"net/http"
)

type BudgetPageData struct {
	PageTitle   string
	BudgetSheet Budget
}

func handle_index(w http.ResponseWriter, req *http.Request) {
	budget := createBudget(1, 2025, 5195.10)
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, BudgetPageData{"Budget time", budget})
}

func main() {
	http.HandleFunc("/", handle_index)

	http.ListenAndServe(":8090", nil)
}
