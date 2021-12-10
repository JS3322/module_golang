package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3, 4, 5}

	// int a slice ascending
	sort.Sort(sort.IntSlice(a))
	fmt.Println(a)
	// int a slice descending
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)

	//sort.Float64Slice(), sort.StringSlice()
}
