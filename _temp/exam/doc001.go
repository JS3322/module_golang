package main

import (
	"fmt"
	"sort"
)

func main() {
	record001 := []string{"t", "c", "r", "o"}
	record002 := []int{5, 6, 7, 8}
	record003 := [][][]int{{{1, 2}, {-1, 2}}, {{4, 5}, {2, 7}}}
	record004 := make(map[int]int)
	//fmt.Println(solution(record001))
	//fmt.Println(solution(record003, record004))
	fmt.Println(record001)
	fmt.Println(record002)
	fmt.Println(record003)
	fmt.Println(record004)
	fmt.Println(sliceSort())
}

func solution(record []string) []int {
	var answer []int

	return answer
}

func array(input [][][]int, map_1 map[int]int) int {
	fmt.Println(input)
	map_1[30] = 500
	map_1[50] = 800
	map_1[20] = 100
	for k, v := range map_1 {
		fmt.Println("array = ", k, v)
	}
	return 5
}

//slice sort
type myDataType struct {
	name string
	age  int
}

func sliceSort() []myDataType {
	mySlice := make([]myDataType, 0)
	mySlice = append(mySlice, myDataType{"js1", 42})
	mySlice = append(mySlice, myDataType{"js2", 28})
	mySlice = append(mySlice, myDataType{"js3", 38})
	fmt.Println(mySlice)
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].age < mySlice[j].age
	})
	return mySlice
}
