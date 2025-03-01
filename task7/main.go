package main

import "fmt"

func mergeChannels(channels ...chan int) chan int {
	resChan := make(chan int, len(channels))

	for _, channel := range channels {
		resChan <- <-channel
		close(channel)
	}

	return resChan
}

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 1
	ch2 := make(chan int, 1)
	ch2 <- 2
	ch3 := make(chan int, 1)
	ch3 <- 3

	mergeChan := mergeChannels(ch1, ch2, ch3)
	close(mergeChannels())

	for val := range mergeChan {
		fmt.Print(val)
	}
}
