package main

import "log"

func fnV(arr [4]int) {
	for i, _ := range arr {
		arr[i] = 0
	}
	log.Println(arr) // [0, 0, 0, 0]
}

func fnP(arr *[4]int) {
	for i, _ := range arr {
		arr[i] = 0
	}
	log.Println(arr) // [0, 0, 0, 0]
}

func main() {
	arr := [4]int{1, 2, 3, 4}
	fnV(arr)
	log.Println(arr) // [1, 2, 3, 4]
	fnP(&arr)
	log.Println(arr) // [0, 0, 0, 0]
}
