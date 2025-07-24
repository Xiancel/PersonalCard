package module

// структура яка містить
type Items struct {
	Id    int    // ID предмету
	Name  string // назва предмета
	Grade int    // оцінка предмета
	Notes string // нотатки предмета
}

// слайсл для зберегання предметів
var ItemList []Items
