package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("test.txt.gz")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	r, err := gzip.NewReader(file)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Close()

	//decompress func is r read > read decompress data
	b, err := ioutil.ReadAll(r)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
