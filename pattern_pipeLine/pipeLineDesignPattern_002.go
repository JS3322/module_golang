// both input and output channels are parameters
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func readSomething(outCh chan<- string) {

	defer close(outCh)

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		outCh <- fmt.Sprintf("line: %d", i)
	}

}

func fetchSomething(inCh <-chan string, outChCh chan<- chan string) {

	defer close(outChCh)

	for line := range inCh {

		outCh := make(chan string)

		go getLine(line, outCh)
		outChCh <- outCh

	}

}

func getLine(line string, outCh chan<- string) {
	defer close(outCh)
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	outCh <- fmt.Sprintf("%s ... fetched!", line)
}

func distribute(inCh <-chan chan string, outCh chan<- string) {
	defer close(outCh)
	for ch := range inCh {
		outCh <- <-ch
	}
}

func printSomething(inCh <-chan string) {
	for line := range inCh {
		fmt.Println(line)
	}
}

func test(intensity int) {

	start := time.Now()

	readCh := make(chan string)
	fetchCh := make(chan chan string, intensity)
	distributeCh := make(chan string)

	go readSomething(readCh)
	go fetchSomething(readCh, fetchCh)
	go distribute(fetchCh, distributeCh)
	printSomething(distributeCh)

	fmt.Println("done", time.Now().Sub(start))

}

func main() {
	test(20)
	test(40)
}
