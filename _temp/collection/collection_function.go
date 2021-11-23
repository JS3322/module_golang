package main

import (
	"fmt"
	"strings"
)

//문자열 t의 첫 위치를 반환하며 매칭이 없을시에는 -1을 반환합니다.
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

//문자열 t가 슬라이스에 존재하면 true를 반환합니다.
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

//슬라이스의 문자열중 하나가 조건 f를 만족하면 true를 반환합니다.
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

//슬라이스의 문자열 모두가 조건 f를 만족하면 true를 반환합니다.
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

//슬라이스에서 조건 f를 만족하는 모든 문자열을 포함하는 새로운 슬라이스를 반환합니다.
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

//기존 슬라이스에 있는 각각의 문자열에 함수 f를 적용한 결괏값들을 포함하는 새로운 슬라이스를 반환합니다.
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
func main() {
	//다양한 컬렉션 함수를 사용해봅시다.
	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))
	//위의 에시 모두 익명함수를 사용했지만, 올바른 타입을 가진 명명된 함수를 사용할 수도 있습니다.
	fmt.Println(Map(strs, strings.ToUpper))
}
