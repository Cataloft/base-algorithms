package main

import "fmt"

//Удалить i-ый элемент из слайса.

// записываем последний элемент в тот который хотим удалить
func deleteReorder(slc []string, idx int) []string {
	slc[idx] = slc[len(slc)-1]
	slc = slc[:len(slc)-1]
	return slc
}

// сдвиг влево
func deleteSlow(slc []string, idx int) []string {
	copy(slc[idx:], slc[idx+1:])
	slc = slc[:len(slc)-1]
	return slc
}

func main() {
	slc := []string{"a", "b", "c", "d", "e", "f"}
	var idx int
	fmt.Scan(&idx)

	fmt.Println("Удаление с изменением порядка элементов", deleteReorder(slc, idx))
	fmt.Println("Удаление без изменения порядка элементов", deleteSlow(slc, idx))
}
