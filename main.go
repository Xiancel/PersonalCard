package main

import (
	"fmt"
	"net/http"
)

func itemListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "empty item list unluck ma boy!")
}
func main() {
	http.HandleFunc("/", itemListHandler)
	fmt.Println("Server Start http://localhost:8080")
	fmt.Println("Маршрути :\n GET / (item list) vse))")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server Caput unlucky")
		return
	}
}
