package main

import (
	"fmt"
)

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 0}
	fmt.Println(solution(input))
}

func solution(numbers []int) int {
	answer := 0
	answerArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range numbers {
		if v != 0 {
			answerArray[v-1] = 0
		}
	}
	for _, v := range answerArray {
		answer += v
	}
	return answer
}
