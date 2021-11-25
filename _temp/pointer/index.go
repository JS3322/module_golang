package pointer

import (
	"fmt"
)

func foo(n *int) {
	*n = 20
}

func main() {
	n := 10
	foo(&n)

	// 20
	fmt.Println(n)
}
