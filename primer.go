package main

import (
	"fmt"
	"os"
	"strconv"
)

func IsPrime(n int) bool {
	switch n {
	case 1, 2, 3:
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	w := 2
	for i*i <= n {
		if n%i == 0 {
			return false
		}
		i += w
		w = 6 - w
	}

	return true
}

func main() {
	args := os.Args
	for i, a := range args {
		if i == 0 {
			continue
		}
		n, err := strconv.Atoi(a)
		if err == nil {
			if IsPrime(n) {
				fmt.Printf("%d is prime.\n", n)
			} else {
				fmt.Printf("%d is not prime.\n", n)
			}
		} else {
			fmt.Printf("%s is not a number.\n", a)
		}
	}
}

// TEST RUN
// go run .\primer.go 2 4 6 7 a 9 c 11 x 12 13
