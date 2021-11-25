package main

import (
	"fmt"
)

func main() {
	fmt.Println(solution(13, 17))
}

func solution(left int, right int) int {
	//true > +
	//false > -
	//map key is number, value is calc
	var answer int
	var count int
	numberMap := map[int]bool{}
	for i := left; i <= right; i++ {
		//numberMap[i] = yansu
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count += 1
			}
		}
		if count%2 == 1 {
			numberMap[i] = false
		} else {
			numberMap[i] = true
		}

		count = 0
	}

	for k, v := range numberMap {
		if v == true {
			answer += k
		} else {
			answer -= k
		}

	}
	return answer
}
