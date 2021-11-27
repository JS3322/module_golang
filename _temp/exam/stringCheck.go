package main

import (
	"fmt"
)

func main() {
	a := "a234"
	// b := "33321"
	//fmt.Println(stringCheckInt(a))
	fmt.Println(stringCheckInt(a))
	// fmt.Println(stringCheckInt2("2234"))
	//fmt.Println(stringTransformCheckInt(b))
}

// func solution(s string) bool {

// 	var answer bool

// 	fmt.Println(stringTransformArray(s))

// 	return answer
// }

// // func arrayCheckString(arrayS string[]) {

// // }

// func stringTransformArray(str string) []string {
// 	var result []string
// 	for i := 0; i < len(str); i++ {
// 		fmt.Println(string(str[i]))
// 		fmt.Println(reflect.TypeOf(string(str[i])))
// 		//result.append(string(str[i]))
// 	}
// 	return result
// }

// func stringTransformCheckInt(s string) bool {
// 	// var i int
// 	var e error
// 	var answer bool
// 	_, e = strconv.Atoi(s)
// 	if e == nil {
// 		answer = true
// 	} else {
// 		answer = false
// 	}
// 	// fmt.Println(">>>>", i)
// 	fmt.Println(">>>>e", e)
// 	return answer
// }

//0~9 ASKII 48~57
// if result[i] == 48 || 49 || 50 || 51 || 52 || 53 || 54 || 55 || 56 || 57 {
func stringCheckInt(s string) bool {
	result := []byte(s)
	var answer bool = false
	if len(s) == 4 || len(s) == 6 {
		answer = false
	} else {
		answer = false
		return answer
	}
	for i := 0; i < len(result); i++ {
		if result[i] == 48 {
			answer = true
			break
		} else if result[i] == 49 {
			answer = true
			break
		} else if result[i] == 50 {
			answer = true
			break
		} else if result[i] == 51 {
			answer = true
			break
		} else if result[i] == 52 {
			answer = true
			break
		} else if result[i] == 53 {
			answer = true
			break
		} else if result[i] == 54 {
			answer = true
			break
		} else if result[i] == 55 {
			answer = true
			break
		} else if result[i] == 56 {
			answer = true
			break
		} else if result[i] == 57 {
			answer = true
			break
		} else {
			break
		}
	}
	return answer
}

// func stringCheckInt2(s string) {
// 	var answer bool
// 	if len(s) != 4 || len(s) != 6 {
// 		answer = true
// 		return answer
// 	}
// 	return answer
// }

import "strconv"
func solution(s string) bool {
    var answer bool
    _, err := strconv.Atoi(s)
if err != nil || len(s) !=4 && len(s) != 6 {
	return answer
}
    answer = true
return true

}
