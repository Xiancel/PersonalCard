package handlers

import (
	"fmt"
	"net/http"
	mod "personalcard/module"
	"strconv"
)

// функція/Хендлер додавання нового предмету
func AddHandler(w http.ResponseWriter, r *http.Request) {
	// перевірка метода
	if r.Method == "POST" {
		// запрос у користувача: Ім'я, оцінки, нотатки для предмета
		name := r.URL.Query().Get("name")
		gradeStr := r.URL.Query().Get("grade")
		notes := r.URL.Query().Get("notes")

		// перевірка на наявність назви
		if name == "" {
			fmt.Println("❌ Помилка: введіть назву предмета")
			return
		}
		// перевод з string в int
		grade, err := strconv.Atoi(gradeStr)
		// перевірка на введеня правельного значення
		if err != nil {
			http.Error(w, "Значення Grade введено неправильно", http.StatusBadRequest)
			fmt.Println("❌ Помилка: Значення Grade введено неправильно")
			return
		}
		// перевірка на діапазон оцінки
		if grade < 0 || grade > 12 {
			fmt.Println("❌ Помилка: Значення Grade поза діапазоном")
			return
		}

		// додавання нового предмета
		newItem := mod.Items{
			Id:    len(mod.ItemList) + 1,
			Name:  name,
			Grade: grade,
			Notes: notes,
		}
		mod.ItemList = append(mod.ItemList, newItem)

		// вивід інформаціЇ
		fmt.Println("✅ ПРЕДМЕТ ДОДАНО УСПІШНО!")
		fmt.Printf("\nID: %d\nНазва: %s\nОцінка: %d/12\nНотатки: %s", newItem.Id, newItem.Name, newItem.Grade, newItem.Notes)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Sorry this method no support")
	}
}
