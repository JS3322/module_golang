package main

import (
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"test.txt.gz",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	w := gzip.NewWriter(file)
	defer w.Close()

	s := "This is a resume written on the basis of skills based on experience. I enjoy researching and developing, including design, backend, architecture and data programming. As a software designer specializing in architects, I want to become a developer who cultivates data operation programming skills. I hope you have the skills to help with your project."
	w.Write([]byte(s))
	// data final save
	w.Flush()
}
