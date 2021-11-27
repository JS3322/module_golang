package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	a := []string{"sun", "bed", "car"}
	b := 1
	fmt.Println(solution(a, b))
	// fmt.Println(arraySort())
}

//배열 참조 후 주소값으로 찾기?
//map key, valuse sort
func solution(strings []string, n int) []string {
	var stringsNum map[string]string
	stringsNum = make(map[string]string)

	var indexString map[int]string
	indexString = make(map[int]string)
	var indexChar map[int]string
	indexChar = make(map[int]string)
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
		fmt.Println("check ", string(strings[i][n]))
		stringsNum[strings[i]] = string(strings[i][n])
		indexChar[i] = string(strings[i][n])
		indexString[i] = strings[i]
	}

	fmt.Println(" >>> ", stringsNum, len(stringsNum), reflect.TypeOf(len(stringsNum)))
	fmt.Println(" !!!! ", indexString)
	fmt.Println(" indexChar ", indexChar)

	arraySort(indexString, indexChar)

	return strings
}

func arraySort(inputString map[int]string, inputChar map[int]string) map[int]string {
	var result map[int]string
	sortValues := make([][]string, 0, len(inputChar))
	for k, v := range inputChar {
		fmt.Println("k", k, v)
		sortValues = append(sortValues, v, string(k))
	}
	sort.Strings(sortValues)

	fmt.Println("sortValue ", sortValues)

	for k, v := range inputString {
		fmt.Println("sort>>", k, v, reflect.TypeOf(k), reflect.TypeOf(v))
		fmt.Println(">>>>", inputChar[k])
	}

	return result
}

func Solution(strings []string, n int) []string {
	sort.Slice(strings, func(i, j int) bool {
		if strings[i][n] == strings[j][n] {
			return strings[i] < strings[j]
		}
		return strings[i][n] < strings[j][n]
	})

	return strings
}
