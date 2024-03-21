package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.CountChar("hellow", 'l'))
	fmt.Println(piscine.CountChar("anne alice okingo", 'n'))
	fmt.Println(piscine.CountChar("  ", ' '))
	fmt.Println(piscine.CountChar("meeeow", 'e'))
	fmt.Println(piscine.CountChar("", 'l'))
	fmt.Println(piscine.CountChar("hellow", 'k'))
}
