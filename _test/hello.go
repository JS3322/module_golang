package main

import (
	"fmt"
)

var f float32 = 11.
var i, j, k int = 1, 2, 3
var stringTest string = "js"
var stringTest2 string
var num int

func main() {
	fmt.Println(stringTest2)
	println(f, j)
	println("hello world")

	fmt.Println("input u a number")
	fmt.Scanln(&num)

	if num > 0 {
		println()
	}
}
