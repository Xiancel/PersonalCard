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

	// перевірка метода
	if r.Method == "GET" {
		// вивід інформації
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
