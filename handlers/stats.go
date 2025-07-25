package handlers

import (
	"fmt"
	"net/http"
	mod "personalcard/module"
)

// функція визначення середніого балу
func avgGrade() float64 {
	var ball float64
	for _, n := range mod.ItemList {
		ball += float64(n.Grade)
	}

	var avg float64 = ball / float64(len(mod.ItemList))

	return avg
}

// функція визначення найкрашого предмета
func bestStats() (int, string) {
	// перевірка на наявність предметів
	if len(mod.ItemList) == 0 {
		return 0, " "
	}
	bestGrade := mod.ItemList[0].Grade
	bestName := mod.ItemList[0].Name

	// визначення найкрашого предмета
	for _, n := range mod.ItemList[1:] {
		if n.Grade > bestGrade {
			bestGrade = n.Grade
			bestName = n.Name
		}
	}
	return bestGrade, bestName
}

// функція визначення найгіршого предмета
func worstStats() (int, string) {
	// перевірка на наявність предметів
	if len(mod.ItemList) == 0 {
		return 0, " "
	}
	worstGrade := mod.ItemList[0].Grade
	worstName := mod.ItemList[0].Name

	// визначення найгіршого предмета
	for _, n := range mod.ItemList[1:] {
		if n.Grade < worstGrade {
			worstGrade = n.Grade
			worstName = n.Name
		}
	}
	return worstGrade, worstName
}

// функція/хендлер для виводу статистки навчання
func StatsHandler(w http.ResponseWriter, r *http.Request) {
	// виклик функції
	bestGrade, bestName := bestStats()
	worstGrade, worstName := worstStats()
	avg := avgGrade()
	total := len(mod.ItemList)
	// перевірка метода
	if r.Method == "GET" {
		// вивід інформації
		fmt.Println("\n📊 === СТАТИСТИКА НАВЧАННЯ ===")
		fmt.Printf("Всього предметів: %d\n", total)
		fmt.Printf("Середній бал: %.2f/12\n", avg)
		fmt.Printf("Найкраща оцінка: %d/12 (%s)\n", bestGrade, bestName)
		fmt.Printf("Найгірша оцінка: %d/12 (%s)\n", worstGrade, worstName)

		// html розмітка
		html := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Статистика навчання</title>
			<style>
				body { font-family: sans-serif; padding: 20px; }
				ul { list-style-type: none; padding: 0; }
				li { margin-bottom: 8px; }
			</style>
		</head>
		<body>
			<h1>СТАТИСТИКА НАВЧАННЯ</h1>
			<ul>
				<li><strong>Всього предметів:</strong> %d</li>
				<li><strong>Середній бал:</strong> %.2f/12</li>
				<li><strong>Найкраща оцінка:</strong> %d/12 (%s)</li>
				<li><strong>Найгірша оцінка:</strong> %d/12 (%s)</li>
			</ul>
		</body>
		</html>
		`, total, avg, bestGrade, bestName, worstGrade, worstName)

		// формування html сторінки
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(html))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Sorry this method no support")
	}
}
