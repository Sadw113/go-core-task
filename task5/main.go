package main

import "fmt"

func main() {
	slice1 := []int{65, 3, 58, 678, 64}
	slice2 := []int{64, 2, 3, 43}

	resSlice, exist := intersectionOfValues(slice1, slice2)
	if exist {
		fmt.Println(exist, resSlice)
	}
}

func intersectionOfValues(slice1, slice2 []int) ([]int, bool) {
	var exist bool
	var resSlice []int
	hashTable := make(map[int]int)

	for _, v := range slice1 {
		hashTable[v]++
	}

	for _, v := range slice2 {
		hashTable[v]++
	}

	for key, val := range hashTable {
		if val > 1 {
			exist = true
			resSlice = append(resSlice, key)
		}
	}

	return resSlice, exist
}
