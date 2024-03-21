package main

import (
	"fmt"
)

func digiLen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	}
	if n < 0 {
		n = -n
	}
	count := 0
	for n > 0 {
		n /= base
		count++

	}

	return count
}
func main() {
	k := digiLen(100, 2)
	fmt.Println(k)
}
