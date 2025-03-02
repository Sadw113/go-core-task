package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	intChan := make(chan uint8)
	floatChan := make(chan float64)
	wg.Add(10)
	go func() {
		for i := range 10 {
			intChan <- uint8(i)
		}
		close(intChan)
	}()
	go func() {
		defer close(floatChan)
		for {
			num, ok := <-intChan
			if !ok {
				return
			}
			res := float64(num)
			floatChan <- res * res * res
		}
	}()
	for val := range floatChan {
		fmt.Println(val)
	}
}
