package main

import (
	"fmt"
	"strconv"
)

func main() {
	book := 300 % 500
	var age, age2 int = 2, 3444

	fmt.Printf("hello,world! %v %d %o\nasd", book, age, age2)

	for i := 1; i <= 5; i++ {
		if i >= 2 {
			fmt.Printf("amd,yes!")

		} else {
			fmt.Printf("animate")
		}
	}

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Printf(strconv.Itoa(i * j))
			NewPrint("all")
		}

	}
}

func NewPrint(num string) int {
	fmt.Printf(num)
	return 1
}
