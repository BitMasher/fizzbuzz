package main

import (
	"os"
	"strconv"
)

const TestSize = 2_000_000

func main() {
	for i := 1; i <= TestSize; i++ {
		if i%15 == 0 {
			os.Stdout.Write([]byte("fizzbuzz"))
		} else if i%3 == 0 {
			os.Stdout.Write([]byte("fizz"))
		} else if i%5 == 0 {
			os.Stdout.Write([]byte("buzz"))
		} else {
			os.Stdout.Write([]byte(strconv.Itoa(i)))
		}
	}
}
