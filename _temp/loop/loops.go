package main

import (
	"fmt"
)

func main() {
	i := 0
	for i = 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}
		if i == 95 {
			break
		}
		fmt.Print(i, " ")
	}
	i = 0
	fmt.Println("-------------------", i)

	anArray := [5]int{0, 1, -1, 2, -2}
	for i, value := range anArray {
		fmt.Println("index:", i, "value:", value)
	}
}
