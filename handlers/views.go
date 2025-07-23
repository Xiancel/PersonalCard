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
		checkId := r.URL.Query().Get("ID")
		if checkId == "" {
			http.Error(w, "ID не вказано", http.StatusBadRequest)
			fmt.Println("Enter Id") //может удалить потом
			return
		}

		inputId, err := strconv.Atoi(checkId)
		if err != nil {
			http.Error(w, "Невірний формат ID", http.StatusBadRequest)
			return
		}

		item := filterById(inputId)
		if item.Id == 0 {
			http.Error(w, "Предмет не знайдено", http.StatusBadRequest)
			return
		}
		fmt.Println("\n=== ПЕРЕГЛЯД ПРЕДМЕТУ ===")
		fmt.Printf("ID: %d\nНазва: %s\nОцінка: %d/12\nНотатки: %s", item.Id, item.Name, item.Grade, item.Notes)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Sorry this method no support")
	}
}
