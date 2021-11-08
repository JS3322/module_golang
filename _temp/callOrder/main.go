package main

import (
	"util"
)

func init() {
	println("main init() 실행")
}

func main() {
	println("main 실행")
	util.Test1()
	util.Test2()
	//util.test() // 호출 불가
}
