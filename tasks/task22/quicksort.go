// Package task22 Написать быструю
// сортировку встроенными методами языка.
package task22

import (
	"fmt"
	"math/rand"
	"sort"
)

// Quicksort ...
func Quicksort(slice []int) []int {
	// Для примера возьмем []int{10, 5, 6, 2, 3, 7}
	// Если в слайсе 1 и менее элемент - он уже отсортирован
	if len(slice) <= 1 {
		return slice
	}
	// Выбираем левую и правую границу
	leftB, rightB := 0, len(slice)-1
	// Выбираем случайный опорный элемент. Пусть 4
	based := rand.Intn(len(slice))
	// Ставим опорный элемент в самый конец
	slice[based], slice[rightB] = slice[rightB], slice[based]
	// slice[4], slice[5] = slice[5], slice[4]
	// Теперь слайс выглядит []int{10, 5, 6, 2, 7, 3}
	// Идем по слайсу
	for i := range slice {
		// Если элемент меньше опорного
		if slice[i] < slice[rightB] {
			// Первый раз случится при i=3
			slice[leftB], slice[i] = slice[i], slice[leftB]
			// []int{2, 5, 6, 10, 7, 3}
			// Поднимаем на единцу левую границу
			leftB++
		}
	}
	// Ставим опорный элемент перед последним меньшим
	slice[leftB], slice[rightB] = slice[rightB], slice[leftB]
	// []int{2, 3, 6, 10, 7, 5}
	Quicksort(slice[:leftB])   // Повторяем для левой части []int{2}
	Quicksort(slice[leftB+1:]) // Повторяем для правой части []int{6, 10, 7, 5}

	return slice
}

// Start ...
func Start() {
	s := []int{-100, 121, 20, 30, 4, 5, -7, 8, 900, 2, 3, 5, 10, 12, -121, 121}
	fmt.Println("Unsorted:", s)
	Quicksort(s)
	fmt.Println("Sorted:", s)
	ss := []int{-100, 121, 20, 30, 4, 5, -7, 8, 900, 2, 3, 5, 10, 12, -121, 121}

	// Слишком тупо, но там тоже квиксорт
	sort.Ints(ss)
	fmt.Println("Sorted std:", ss)

}
