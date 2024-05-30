package main

import (
	"os"
	"strconv"
	"strings"
)

const TestSize = 2_000_000

func main() {
	cache := make([]string, TestSize)
	for i := 1; i <= TestSize; i++ {
		if i%15 == 0 {
			cache[i-1] = "fizzbuzz"
		} else if i%3 == 0 {
			cache[i-1] = "fizz"
		} else if i%5 == 0 {
			cache[i-1] = "buzz"
		} else {
			cache[i-1] = strconv.Itoa(i)
		}
	}
	buff := strings.Join(cache[:], "\n")
	os.Stdout.Write([]byte(buff))
}
