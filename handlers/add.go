package handlers

import (
	"fmt"
	"net/http"
	mod "personalcard/module"
	"strconv"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.URL.Query().Get("name")
		gradeStr := r.URL.Query().Get("grade")
		notes := r.URL.Query().Get("notes")

		if name == "" {
			fmt.Println("Enter name pls")
			return
		}
		grade, err := strconv.Atoi(gradeStr)
		if err != nil {
			http.Error(w, "Значення Grade введено неправильно", http.StatusBadRequest)
			return
		}

		newItem := mod.Items{
			Id:    len(mod.ItemList) + 1,
			Name:  name,
			Grade: grade,
			Notes: notes,
		}

		mod.ItemList = append(mod.ItemList, newItem)

		fmt.Println("✅ ПРЕДМЕТ ДОДАНО УСПІШНО!")
		fmt.Printf("\nID: %d\nНазва: %s\nОцінка: %d/12\nНотатки: %s", newItem.Id, newItem.Name, newItem.Grade, newItem.Notes)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Sorry this method no support")
	}
}
