package handlers

import (
	"fmt"
	"net/http"
	mod "personalcard/module"
	"strconv"
)

// функція для пошуку предмата за його ID
func searchById(inputId int) mod.Items {
	for _, item := range mod.ItemList {
		if item.Id == inputId {
			return item
		}
	}
	return mod.Items{}
}

// функція/Хендлер для перегляду предмета за його ID
func ViewsHandler(w http.ResponseWriter, r *http.Request) {
	// перевірка метода
	if r.Method == "GET" {
		// запрос ID у користувача
		checkId := r.URL.Query().Get("id")

		// перевірка на вказання ID
		if checkId == "" {
			http.Error(w, "ID не вказано", http.StatusBadRequest)
			fmt.Println("❌ ПОМИЛКА: ID параметр обов'язковий")
			return
		}

		// перетворення id з string в int
		inputId, err := strconv.Atoi(checkId)
		// перевірка на правильність формату
		if err != nil {
			http.Error(w, "Невірний формат ID", http.StatusBadRequest)
			return
		}

		// виклик функції пошуку предмета
		item := searchById(inputId)
		// перевірка на существование предмета з таким ID
		if item.Id == 0 {
			http.Error(w, "Not Found", http.StatusBadRequest)
			fmt.Printf("❌ ПОМИЛКА: Предмет з ID=%d не знайдено\n", inputId)
			return
		}

		// вивід інформації про знайдений предмет
		fmt.Println("\n=== ПЕРЕГЛЯД ПРЕДМЕТУ ===")
		fmt.Printf("ID: %d\nНазва: %s\nОцінка: %d/12\nНотатки: %s", item.Id, item.Name, item.Grade, item.Notes)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Sorry this method no support")
	}
}
