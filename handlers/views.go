package handlers

import (
	"fmt"
	"net/http"
	mod "personalcard/module"
	"strconv"
)

func filterById(inputId int) mod.Items {
	for _, item := range mod.ItemList {
		if item.Id == inputId {
			return item
		}
	}
	return mod.Items{}
}
func ViewsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		checkId := r.URL.Query().Get("id")
		if checkId == "" {
			http.Error(w, "ID не вказано", http.StatusBadRequest)
			fmt.Println("❌ ПОМИЛКА: ID параметр обов'язковий")
			return
		}

		inputId, err := strconv.Atoi(checkId)
		if err != nil {
			http.Error(w, "Невірний формат ID", http.StatusBadRequest)
			return
		}

		item := filterById(inputId)
		if item.Id == 0 {
			http.Error(w, "Not Found", http.StatusBadRequest)
			fmt.Printf("❌ ПОМИЛКА: Предмет з ID=%d не знайдено\n", inputId)
			return
		}
		fmt.Println("\n=== ПЕРЕГЛЯД ПРЕДМЕТУ ===")
		fmt.Printf("ID: %d\nНазва: %s\nОцінка: %d/12\nНотатки: %s", item.Id, item.Name, item.Grade, item.Notes)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Sorry this method no support")
	}
}
