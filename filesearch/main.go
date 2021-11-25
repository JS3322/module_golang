// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// )

// func main() {

// 	targetDir := "/test/dir"
// 	files, err := ioutil.ReadDir(targetDir)
// 	if err != nil {
// 		return err
// 	}
// 	for _, file := range files {
// 		// 파일명
// 		fmt.Println(file.Name())
// 		// 파일의 절대경로
// 		fmt.Println(fmt.Sprintf("%v/%v", targetDir, file.Name()))
// 	}
// }
