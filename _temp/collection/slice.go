package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3}
	a[1] = 10
	fmt.Println(a)

	s := make([]int, 5, 10)
	//len : slice length, cap: capacity
	println(len(s), cap(s))

	//nil slice
	var t []int
	if t == nil {
		println("nil t")
	}
	println(len(t), cap(t))

	//sub slice
	ss := []int{1, 2, 3, 4, 5, 6}
	ss = ss[2:5]
	//print array address and cap[3/4]
	println(ss)
	//fmt print valuse
	fmt.Println(ss)

	//clice append copy
	s_apend := []int{0, 1}
	s_apend = append(s_apend, 2, 4)
	fmt.Println(s_apend)

	//slice is [pointer, length, cap] valuse > sub slice is length and max cap valuse change

}
