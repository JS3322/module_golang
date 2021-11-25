// defer > finaly
// panic > function close
// recover > panic > nomalcy
package main

//defer example
// func main() {
// 	f, err := os.Open("test.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	bytes := make([]byte, 1024)
// 	f.Read(bytes)
// 	println(len(bytes))
// }

//panic example
// func main() {
// 	openFile("test.txt")
// 	println("panic function > no message")
// }
// func openFile(fn string) {
// 	f, err := os.Open(fn)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()
// 	defer println("close")
// }

// //recover example
// func main() {
//     openFile("test.txt")
//     println("Done")
// }

// func openFile(fn string) {
//     defer func() {
//         if r := recover(); r != nil {
//             fmt.Println("OPEN ERROR", r)
//         }
//     }()

//     f, err := os.Open(fn)
//     if err != nil {
//         panic(err)
//     }

//     defer f.Close()
// }
