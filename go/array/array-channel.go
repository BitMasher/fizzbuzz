package main

import (
	"fmt"
	"strconv"
)

const TestSize = 2_000_000

func proc(c *[2_000_000]string, pipe chan int) {
	for {
		i, more := <-pipe
		if more {
			if i%15 == 0 {
				c[i-1] = "fizzbuzz"
			} else if i%3 == 0 {
				c[i-1] = "fizz"
			} else if i%5 == 0 {
				c[i-1] = "buzz"
			} else {
				c[i-1] = strconv.Itoa(i)
			}
		} else {
			return
		}
	}
}

func main() {
	var cache [2_000_000]string
	pipe := make(chan int, 8)
	go proc(&cache, pipe)
	go proc(&cache, pipe)
	go proc(&cache, pipe)
	go proc(&cache, pipe)
	go proc(&cache, pipe)
	go proc(&cache, pipe)
	go proc(&cache, pipe)
	go proc(&cache, pipe)
	for i := 1; i <= TestSize; i++ {
		pipe <- i
	}
	close(pipe)
	fmt.Println(cache)
}
