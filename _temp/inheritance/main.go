package main

import (
	overrideFolder "module_golang/_temp/inheritance/override"
)

func main() {
	// human := overrideFolder.NewHuman("alice")
	// fmt.Println(human.GetName())
	// human.SetName("bob")
	// fmt.Println(human.GetName())

	// man := overrideFolder.NewMan("bob")
	// man.Check() // Human의 메소드를 자신이 구현한 것처럼 불러낼 수 있다.
	human := overrideFolder.NewHuman("alice")
	woman := overrideFolder.NewWoman("emily")

	// Human도 Woman도 Speaker 인터페이스로 다룰 수 있다.
	speak(human)
	speak(woman)
}

func speak(speaker overrideFolder.Speaker) {
	speaker.Speak()
}
