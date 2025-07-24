package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	mod "personalcard/module"
)

// функція/Хендлер
func ItemListHandler(w http.ResponseWriter, r *http.Request) {
	// перевірка методу
	if r.Method == "GET" {
		// html розмітка
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Список предметів</title>
			<style>
				body { font-family: sans-serif; padding: 20px; }
				ul { list-style-type: none; padding: 0; }
				li { margin-bottom: 8px; }
			</style>
		</head>
		<body>
			<h1>Список предметів</h1>
			<ul>
				{{range .}}
					<li>{{.Id}}. {{.Name}} - Оцінка: {{.Grade}}/12{{.Notes}}</li>
				{{end}}
			</ul>
			<p><strong>Всього предметів: {{len .}}</strong></p>
		</body>
		</html>
		`
		// формування html сторінки
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		t := template.Must(template.New("items").Parse(html))
		t.Execute(w, mod.ItemList)

		// вивід всіх предметів користувачу
		fmt.Println("\n === МОЯ КАРТОТЕКА ПРЕДМЕТІВ ===")
		fmt.Printf("\nПредмети (%d):\n", len(mod.ItemList))
		for i, n := range mod.ItemList {
			fmt.Printf("%d. %s - Оцінка: %d/12\n", i+1, n.Name, n.Grade)
		}
		fmt.Println("\nВсього предметів:", len(mod.ItemList))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Sorry this method no support")
	}
}
