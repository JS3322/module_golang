package main

import (
	"sort"
	"strings"
)

func solution(s string) {
	test2 := make([]string, 3)
	test2[0] = "test1"
	test2[1] = "test2"
	test2[2] = "test3"

	split := strings.Split(s, "")
	sort.Sort(sort.Reverse(sort.StringSlice(split)))
	check := ""
	for _, str := range split {
		check += str
	}

	println(check)
	//return check

}

func main() {
	solution("aVddds")

}

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	strs := []string{"c", "a", "b"}
// 	sort.Strings(strs)
// 	fmt.Println("Strings:", strs)

// 	//sort int = sort.Ints(INT)
// 	//sorted check = sort.IntsAreSorted(ints)

// }

// Custom sort
// type ByLength []string

// func (s ByLength) Len() int {
// 	return len(s)
// }
// func (s ByLength) Swap(i, j int) {
// 	s[i], s[j] = s[j], s[i]
// }
// func (s ByLength) Less(i, j int) bool {
// 	return len(s[i]) < len(s[j])
// }
// func main() {
// 	fruits := []string{"peach", "banana", "kiwi"}
// 	sort.Sort(ByLength(fruits))
// 	fmt.Println(fruits)
// }
