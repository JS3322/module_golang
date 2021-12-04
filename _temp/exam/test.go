package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("hello")

	// //for
	// for i := 0; i <= MAX; i++ {
	// 	fmt.Println("check", i)
	// }

	//range
	threeD := [][][]int{{{1, 2}, {-1, 2}}, {{4, 5}, {2, 7}}}
	for _, v := range threeD {
		for _, m := range v {
			for _, s := range m {
				fmt.Println(s)
			}
		}
	}
	rageTest := []int{1, 2, 3, 4}
	rageTest = append(rageTest, 7)
	for index, v := range rageTest {
		fmt.Println("check", index, v)
	}

	//rnage map
	rageTestMap := make(map[string]int)
	rageTestMap["111"] = 123
	rageTestMap["222"] = 223344
	for k, v := range rageTestMap {
		fmt.Println("check", k, v)
	}

	//delete map
	fmt.Println(rageTestMap)
	delete(rageTestMap, "222")
	fmt.Println(rageTestMap)

	// //map key search
	_, ok := rageTestMap["111"]
	if ok {
		fmt.Println("true!!!")
	} else {
		fmt.Println("false!!!")
	}

	//reslicing 참조 값이므로 원본 변경 유의
	si := []int{1, 2, 3, 4, 5}
	resi := si[1:3]
	resiAll := si[:]
	resi[0] = -2
	resi[1] = -3
	fmt.Println(si)
	fmt.Println(resiAll)

	//byte slice
	by := []byte("12341342ffsdfs")
	fmt.Println(by)

	//copy
	a6 := []int{3, 5, 7, 8, 2, 5}
	a6re := make([]int, 6)
	copy(a6re, a6)
	fmt.Println("copy", a6re)

	//struct sort
	type sortExam struct {
		index   int
		sortNum int
	}

	testSlice := make([]sortExam, 0)
	testSlice = append(testSlice, sortExam{0, 1231})
	testSlice = append(testSlice, sortExam{1, 231})
	testSlice = append(testSlice, sortExam{2, 4231})

	sort.Slice(testSlice, func(i, j int) bool {
		return testSlice[i].sortNum < testSlice[j].sortNum
	})
	fmt.Println(testSlice)

	//sort
	strings := []string{}
	strings = append(strings, "czz")
	strings = append(strings, "fes")
	strings = append(strings, "zzz")
	fmt.Println(Solution(strings, 1))

	//err
	// f, err := strings["error"]
	// if err != nil {
	// 	fmt.Println("erre!!!!")
	// }

	//slice in array
	s200 := []int{1, 2, 3}
	a200 := [3]int{4, 5, 6}
	ref200 := a200[:]
	t200 := append(s200, ref200...)
	fmt.Println("t200", t200)

}

func Solution(strings []string, n int) []string {
	sort.Slice(strings, func(i, j int) bool {
		if strings[i][n] == strings[j][n] {
			return strings[i] < strings[j]
		}
		return strings[i][n] < strings[j][n]
	})

	return strings
}
