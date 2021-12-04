package main

import (
	"fmt"
)

func main() {
	input := []string{"#*<*###", "#*<*<*#", "*#*><*#", "#<#>>##", ">#*#*#*", "###*##*"}
	fmt.Println(solution(input))
}

func sort_array(drum []string) []int {
	result := []int{}
	for _, str := range drum {
		for _, s := range str {
			result = append(result, int(s))
		}
	}
	return result
}

func solution(drum []string) int {
	count := 0
	result := 0
	//change int array
	result_cehck := sort_array(drum)
	//drum len check
	var len_drum int
	len_drum = len(drum)
	var len_drum_array int
	for _, str := range drum {
		len_drum_array = len(str)
	}
	//temp var

	var i int
	var j int
	//#><* check
	for index_array := 0; index_array < len_drum_array; index_array++ {
		j = index_array
		for i = index_array; i < len(result_cehck); i += len_drum {
			fmt.Println(result_cehck[i])
			if result_cehck[i] == 62 {
				i += 1
				i -= len_drum
			} else if result_cehck[i] == 60 {
				i -= 1
				i -= len_drum
			} else if result_cehck[i] == 42 {
				count++
				if count >= 2 {
					result++
					break
				}
			}
		}
		count = 0
		index_array = j
		fmt.Println("____")
	}
	return result
}

// # down 35
// > right 62
// < left
// * 1. down 2. hold 42
