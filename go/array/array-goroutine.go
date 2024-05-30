package main

import (
	"os"
	"strconv"
	"strings"
)

const TestSize = 2_000_000

func proc(s int, e int, c *[2_000_000]string, done chan bool) {
	for i := s; i <= e; i++ {
		if i%15 == 0 {
			c[i-1] = "fizzbuzz"
		} else if i%3 == 0 {
			c[i-1] = "fizz"
		} else if i%5 == 0 {
			c[i-1] = "buzz"
		} else {
			c[i-1] = strconv.Itoa(i)
		}
	}
	done <- true
}

const threads = 1

func main() {
	var cache [2_000_000]string
	inc := TestSize / threads
	done := make(chan bool, threads)
	for i := 0; i < threads; i++ {
		go proc((i*inc)+1, (i+1)*inc, &cache, done)
	}
	for i := 0; i < threads; i++ {
		<-done
	}
	buff := strings.Join(cache[:], "\n")
	os.Stdout.Write([]byte(buff))
}
