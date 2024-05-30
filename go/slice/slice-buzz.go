package main

import (
	"os"
	"strconv"
	"strings"
)

const TestSize = 2_000_000

func main() {
	cache := make([]string, 0)
	for i := 1; i <= TestSize; i++ {
		if i%15 == 0 {
			cache = append(cache, "fizzbuzz")
		} else if i%3 == 0 {
			cache = append(cache, "fizz")
		} else if i%5 == 0 {
			cache = append(cache, "buzz")
		} else {
			cache = append(cache, strconv.Itoa(i))
		}
	}
	buff := strings.Join(cache[:], "\n")
	os.Stdout.Write([]byte(buff))
}
