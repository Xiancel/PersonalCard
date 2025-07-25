package main

import (
	"fmt"
	"net/http"
	"personalcard/halper"
	"personalcard/handlers"
)

// головна функція
func main() {
	// додавання хендлерів
	http.HandleFunc("/", handlers.ItemListHandler)
	http.HandleFunc("/stats", handlers.StatsHandler)
	http.HandleFunc("/add", handlers.AddHandler)
	http.HandleFunc("/views", handlers.ViewsHandler)
	fmt.Println("🚀 Server Start http://localhost:8080")

	// вивід маршрутів
	halper.Routes()

	// ініціалізація сервара
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server Caput unlucky")
		return
	}
}
