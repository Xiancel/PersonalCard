package main

import (
	"fmt"
	"html/template"
	"net/http"
	mod "personalcard/module"
)

func itemListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
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
					<li>{{.Id}}. {{.Name}} - Оцінка: {{.Grade}}/{{.Notes}}</li>
				{{end}}
			</ul>
			<p><strong>Всього предметів: {{len .}}</strong></p>
		</body>
		</html>
		`
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		t := template.Must(template.New("items").Parse(html))
		t.Execute(w, mod.ItemList)

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
func avgGrade() float64 {
	var ball float64
	for _, n := range mod.ItemList {
		ball += float64(n.Grade)
	}

	var avg float64 = ball / float64(len(mod.ItemList))

	return avg
}

func bestStats() (int, string) {
	if len(mod.ItemList) == 0 {
		return 0, " "
	}
	var bestGrade int
	var bestName string

	bestGrade = mod.ItemList[0].Grade
	bestName = mod.ItemList[0].Name

	for _, n := range mod.ItemList[1:] {
		if n.Grade > bestGrade {
			bestGrade = n.Grade
			bestName = n.Name
		}
	}
	return bestGrade, bestName
}

func worstStats() (int, string) {
	if len(mod.ItemList) == 0 {
		return 0, " "
	}
	var worstGrade int
	var worstName string

	worstGrade = mod.ItemList[0].Grade
	worstName = mod.ItemList[0].Name

	for _, n := range mod.ItemList[1:] {
		if n.Grade < worstGrade {
			worstGrade = n.Grade
			worstName = n.Name
		}
	}
	return worstGrade, worstName
}
func statsHandler(w http.ResponseWriter, r *http.Request) {
	bestGrade, bestName := bestStats()
	worstGrade, worstName := worstStats()
	if r.Method == "GET" {
		fmt.Println("=== СТАТИСТИКА НАВЧАННЯ ===")
		fmt.Printf("Всього предметів: %d\n", len(mod.ItemList))
		fmt.Printf("Середній бал: %.2f/12\n", avgGrade())
		fmt.Printf("Найкраща оцінка: %d/12 (%s)\n", bestGrade, bestName)
		fmt.Printf("Найгірша оцінка: %d/12 (%s)\n", worstGrade, worstName)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Sorry this method no support")
	}
}
func routes() {
	fmt.Println("Маршрути:")
	fmt.Println("GET / (item list)")
	fmt.Println("GET /stats (stats)")
}
func main() {
	http.HandleFunc("/", itemListHandler)
	http.HandleFunc("/stats", statsHandler)
	fmt.Println("Server Start http://localhost:8080")
	routes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server Caput unlucky")
		return
	}
}
