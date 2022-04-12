package main

import (
	"fmt"
)

// 보통 function, 이 function은 중첩 될 수 없다
func standardFunction(a1 string, a2 string) string {
	return fmt.Sprintf("%s:%s", a1, a2)
}

func main() {
	// 변수에 할당 된 function
	var assignedFunction1 = standardFunction

	// 변수에 할당되고 중첩 된 익명 function
	var assignedFunction2 = func(arg1 string, arg2 string) string {
		return fmt.Sprintf("%s:%s", arg1, arg2)
	}

	// function을 인수로 받아들이고 function를 반환하는 고차 function
	var functionAsArgumentAndReturn = func(addFn func(int, int) int, arg1 int, arg2 int) func(int) int {
		var out = addFn(arg1, arg2)
		// 클로저를 반환
		return func(numArg int) int {
			return out + numArg
		}
	}

	var out = functionAsArgumentAndReturn(
		func(a, b int) int {
			return a + b
		},
		5,
		10,
	)(10)
	fmt.Println(out) // prints 25

	// 중첩 된 익명 functions
	var nested = func() {
		fmt.Println("outer fn")
		var nested2 = func() {
			fmt.Println("inner fn")
			var nested3 = func() {
				fmt.Println("inner arrow")
			}
			nested3()
		}
		nested2()
	}

	nested() // prints:
	// outer fn
	// inner fn
	// inner arrow

	// function를 반환하는 고차 function
	var add = func(x int) func(y int) int {
		// function은 클로저로 반환된다.
		// 변수 x는 이 방법의 외부 스코프에서 얻어지고 클로저에 기억된다.
		return func(y int) int {
			return x + y
		}
	}

	// 더 많은 커링을 만들기 위해 add 메소드를 사용하고 있다.
	var add10 = add(10)
	var add20 = add(20)
	var add30 = add(30)

	fmt.Println(add10(5)) // 15
	fmt.Println(add20(5)) // 25
	fmt.Println(add30(5)) // 35

	// 즉시 호출 된 익명 function (IIFE)
	(func() {
		fmt.Println("anonymous fn")
	})()
	// prints: anonymous fn

	assignedFunction1("a", "b")
	assignedFunction2("a", "b")
}
