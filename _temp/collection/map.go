//map is type reference > nil map
//var idMap map[int]string

//make의 첫 번째 파라미터로 map 키워드와 값타입을 지정
//make함수는 해시테이블 자료구조를 메모리에 생성하고 메모리를 가리키는 map 값을 리턴
//map 값은 내부적으로 runtime.hmap 구조체 가리키는 포인터 > 이 변수는 해시테이블을 가리키는 map을 가리킴

package main

import "fmt"

func main() {
	var m map[int]string
	m = make(map[int]string)
	m[110] = "app"
	m[211] = "web"
	m[323] = "server"

	n := map[string]string{
		"ip":      "ip port",
		"gateway": "gateway in ip check",
		"hdd":     "ssd change",
	}

	str := m[211]
	println(str)

	noData := m[310]
	println(noData)

	delete(m, 211)
	//str에 valuse reference 이후
	println(str)

	val, exists := n["gateway"]
	if exists {
		println("check out")
		println(val)
	}

	for key, val := range n {
		fmt.Println(key, val)
	}
}
