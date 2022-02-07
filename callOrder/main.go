package main

import (
	util "module_golang/_temp/callOrder/util"
)

func init() {
	println("main init() 실행")
}

func main() {
	println("main 실행")
	util.Test1()
	util.Test2()
}
