package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "one4seveneight"
	fmt.Println(solution(input))
}

func solution(s string) int {

	// eglishNumber := map[string]string{}
	// eglishNumber["zero"] = "0"
	// eglishNumber["one"] = "1"
	// eglishNumber["two"] = "2"
	// eglishNumber["three"] = "3"
	// eglishNumber["four"] = "4"
	// eglishNumber["five"] = "5"
	// eglishNumber["six"] = "6"
	// eglishNumber["seven"] = "7"
	// eglishNumber["eight"] = "8"
	// eglishNumber["nine"] = "9"

	// var val string
	s = strings.Replace(s, "zero", "0", -1)
	s = strings.Replace(s, "one", "1", -1)
	s = strings.Replace(s, "two", "2", -1)
	s = strings.Replace(s, "three", "3", -1)
	s = strings.Replace(s, "four", "4", -1)
	s = strings.Replace(s, "five", "5", -1)
	s = strings.Replace(s, "six", "6", -1)
	s = strings.Replace(s, "seven", "7", -1)
	s = strings.Replace(s, "eight", "8", -1)
	s = strings.Replace(s, "nine", "9", -1)

	answer, _ := strconv.Atoi(s)

	// for k, v := range eglishNumber {
	// 	fmt.Println(k, v)
	// 	val = strings.Replace(s, "one", "noooos", 1)
	// 	fmt.Println(val)

	// }

	// var t string
	// t = strings.Replace(s, "one", "noooos", 1)
	// fmt.Println(t)

	// r := strings.NewReplacer("zero", "0", "one", "1")
	// r.Replace(s)
	// fmt.Println(s)

	//var answer int
	//answer, err = strconv.Atoi(s)

	// for k, v := range eglishNumber {
	// 	// fmt.Println(k, v)
	// 	// strings.Replace(s, k, v, -1)
	// 	// fmt.Println(s)
	// 	s = strings.NewReplacer(k, v)
	// }

	// //b, _ := strconv.Atoi("100")
	// fmt.Println(s)
	// a := "one4seveneight"
	// strings.Replace(a, "one", "1", 1)
	// fmt.Println(a)
	// b := "Hello Hello"
	// strings.Replace(b, "llo", "Go", 2)
	// fmt.Println(b)
	return answer
}

// import (
//     "strconv"
//     "strings"
// )

// func solution(s string) int {
//     m := map[string]string{
//         "zero":  "0",
//         "one":   "1",
//         "two":   "2",
//         "three": "3",
//         "four":  "4",
//         "five":  "5",
//         "six":   "6",
//         "seven": "7",
//         "eight": "8",
//         "nine":  "9",
//     }
//     result := s
//     for word, number := range m {
//         result = strings.ReplaceAll(result, word, number)
//     }
//     num, _ := strconv.Atoi(result)
//     return num
//		}
