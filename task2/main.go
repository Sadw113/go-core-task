package main

import (
	"fmt"
	"math/rand"
)

func main() {
	originalSlice := make([]int, 10, 20)
	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100)
	}

	fmt.Println("Оригинал: ", originalSlice)
	fmt.Println("Все чётные: ", sliceExample(originalSlice))
	copySlice := copySlice(originalSlice)
	fmt.Println("Оригинал: ", originalSlice, " Копия: ", copySlice)
	fmt.Println("Добавление числа 120 к копии")
	copySlice = addElements(copySlice, 120)
	fmt.Println("Оригинал: ", originalSlice, " Копия: ", copySlice)
	fmt.Println("Убираем 2-ой элемент из оригинального массива")
	originalSlice = removeElement(originalSlice, 1)
	fmt.Println("Оригинал: ", originalSlice, " Копия: ", copySlice)
}

func sliceExample(slice []int) []int {
	var resSlice []int
	for _, v := range slice {
		if v%2 == 0 {
			resSlice = append(resSlice, v)
		}
	}

	return resSlice
}

func addElements(slice []int, num int) []int {
	slice = append(slice, num)

	return slice
}

func copySlice(slice []int) []int {
	resSlice := make([]int, len(slice))
	copy(resSlice, slice)

	return resSlice
}

func removeElement(slice []int, idx int) []int {
	var resSlice []int
	if idx < 0 || idx > len(slice) {
		return slice
	}
	left := slice[:idx]
	right := slice[idx+1:]

	resSlice = append(resSlice, left...)
	resSlice = append(resSlice, right...)

	return resSlice
}
