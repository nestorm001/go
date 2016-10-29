package main

import "os"
import (
	"text/template"
)
//In a template, a range action may declare one or two variables.
// If there is only one variable, it is assigned the element (not the index).
// As the documentation states :
//
//this is opposite to the convention in Go range clauses.
const t =
	`{{range $i := .nephews}}Hello {{$i}}
	{{end}}`

func main() {
	nephews := []string{"Huey", "Dewey", "Louie"}
	data := map[string]interface{}{
		"nephews": nephews,
	}
	template.Must(template.New("").Parse(t)).Execute(os.Stdout, data)
}