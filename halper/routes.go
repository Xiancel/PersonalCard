package halper

import "fmt"

// функція для виводу інформації про маршрути в консоль
func Routes() {
	fmt.Println("📌 Маршрути:")
	fmt.Println(" • GET / (item list)")
	fmt.Println(" • GET /stats (stats)")
	fmt.Println(" • POST /add (add new item)")
	fmt.Println(" • GET /views (check item)")
	fmt.Println("⏹️  Для зупинки натисніть Ctrl+C")
}
