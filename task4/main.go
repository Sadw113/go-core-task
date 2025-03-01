package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	resSlice := intersectionOfValues(slice1, slice2)
	fmt.Println(resSlice)
}

func intersectionOfValues(slice1, slice2 []string) []string {
	var resSlice []string
	hashTable := make(map[string]int)

	for _, v := range slice1 {
		hashTable[v]++
	}

	for _, v := range slice2 {
		hashTable[v]++
	}

	for key, val := range hashTable {
		if val > 1 {
			resSlice = append(resSlice, key)
		}
	}

	return resSlice
}
