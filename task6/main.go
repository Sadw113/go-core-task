package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randChan := make(chan int)

	go func() {
		for range 10 {
			source := rand.NewSource(time.Now().Unix())
			rng := rand.New(source)
			res := rng.Intn(100)
			randChan <- res
		}
	}()

	for range 10 {
		val := <-randChan
		fmt.Println(val)
	}

	close(randChan)
}
