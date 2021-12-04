package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	record := []string{"P 300 6", "P 500 3", "S 1000 4", "P 600 2", "S 1200 1"}
	fmt.Println(solution(record))
}

type example struct {
	index      int
	activeType string
	price      int
	count      int
}

func solution(record []string) []int {

	var answer []int
	//calcType flase=after true=before
	var calcType bool
	calcType = false
	pointers := make(map[int]int)

	transExample := make([]example, 0)

	for i := 0; i < len(record); i++ {
		inputArray := stringTransStringArray(record[i])
		tr1, _ := strconv.Atoi(inputArray[1])
		tr2, _ := strconv.Atoi(inputArray[2])
		transExample = append(transExample, example{i, inputArray[0], tr1, tr2})
		//pointer add
		if inputArray[0] != "P" {
			pointers[i] = tr2
		}
	}

	//pointer
	afterWon := purchasePoint(transExample, calcType, pointers)
	calcType = true
	beforeWon := purchasePoint(transExample, calcType, pointers)

	answer = append(answer, beforeWon)
	answer = append(answer, afterWon)

	return answer
}

func stringTransStringArray(s string) []string {
	result := strings.Split(s, " ")
	return result
}

func purchasePoint(transExample []example, calcType bool, pointers map[int]int) int {

	counting := 0
	won := 0
	if calcType {
		// fmt.Println("true purchasePoint // ")

	} else {
		// fmt.Println("false purchasePoint // ", transExample, pointers)
		for k, v := range pointers {
			// fmt.Println("transExample ", transExample[k-1].index, transExample[k-1].count)
			counting = transExample[k-1].count - v
			// fmt.Println("RESULT ", counting)
			if counting >= 0 {
				//변경사항 저장
				// fmt.Println("before ", transExample)
				// fmt.Println("commit BEFORE", transExample[k-1].count, counting)
				won = won + transExample[k-1].price*counting
				transExample[k-1].count = counting
				// fmt.Println("commit AFTER", transExample[k-1].count, counting)
				// fmt.Println("after ", transExample)
			} else if counting < 0 {
				//변경사항 저장
				won = won + transExample[k-1].count*transExample[k-1].price
				transExample[k-1].count = 0
				//윗단계 확인
				won = won + transExample[k-1].price*counting
				transExample[k-2].count = transExample[k-2].count + counting
				// fmt.Println("after <0 ", transExample)

			}
		}
	}

	// fmt.Println("won ", won)
	return won
}
