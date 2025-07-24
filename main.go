package main

import (
	"fmt"
	"net/http"
	"personalcard/handlers"
)

// функція для виводу інформації про маршрути в консоль
func routes() {
	fmt.Println("Маршрути:")
	fmt.Println("GET / (item list)")
	fmt.Println("GET /stats (stats)")
	fmt.Println("POST /add (add new item)")
	fmt.Println("GET /views (check item)")
	fmt.Println("⏹️ Для зупинки натисніть Ctrl+C")
}

// головна функція
func main() {
	// додавання хендлерів
	http.HandleFunc("/", handlers.ItemListHandler)
	http.HandleFunc("/stats", handlers.StatsHandler)
	http.HandleFunc("/add", handlers.AddHandler)
	http.HandleFunc("/views", handlers.ViewsHandler)
	fmt.Println("Server Start http://localhost:8080")

	// вивід маршрутів
	routes()

	// ініціалізація сервара
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server Caput unlucky")
		return
	}
}
