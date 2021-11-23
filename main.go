// package main

// import (
// 	"util"
// )

// func init() {
// 	println("main init() 실행")
// }

// func main() {
// 	println("main 실행")
// 	util.Test1()
// 	util.Test2()
// 	//util.test() // 호출 불가
// }

package main

import (
	"fmt"
	"module_golang/classTest"
)

func main() {
	human := classTest.NewHuman("alice")
	fmt.Println(human.GetName())
	human.SetName("bob")
	fmt.Println(human.GetName())

	man := classTest.NewMan("bob")
	man.Check() // Human의 메소드를 자신이 구현한 것처럼 불러낼 수 있다.
}

// func main() {
// 	println("hello")
// 	music.GetMusic("Adele")
// }
