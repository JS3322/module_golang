package main

import (
	"errors"
	"fmt"
)

// func YourLevel(point int) (int, error) {
//     if point < 0 {
//         return 0, errors.New("Level: 레벨 값은 0보다 커야 합니다.")
//     }
//     if point > 255 {
//         return 0, errors.New("Level: 레벨 값은 255보다 작아야 합니다.")
//     }
//     return point / 10, nil
// }

// func main() {
//     level, err := YourLevel(25)
//     if err != nil {
//         fmt.Println("Error ", err.Error())
//         os.Exit(1)
//     }
//     fmt.Printf("당신의 레벨은 %d 입니다.\n", level)
// }

var StatusPointUnderZero = errors.New("Level: 레벨 값은 0보다 커야 합니다.")
var StatusPointOverflow = errors.New("Level: 레벨 값은 255보다 작아야 합니다.")

func YourLevel(point int) (int, error) {
	if point < 0 {
		return 0, StatusPointUnberZero
	}
	if point > 255 {
		return 0, StatusPointOverflow
	}
	return point / 10, nil
}
func main() {
	level, err := YourLevel(25)
	switch err {
	case StatusPointUnberZero:
		// 에러처리 코드
	case StatusPointOverflow:
		// 에러처리코드
	default:
		// 에러처리 코드
	}
	fmt.Printf("당신의 레벨은 %d 입니다.\n", level)
}

// Positive returns true if the number is positive, false if it is negative.
// func Positive(n int) bool {
// 	return n > -1
// }

// func Check(n int) {
// 	if Positive(n) {
// 			fmt.Println(n, "is positive")
// 	} else {
// 			fmt.Println(n, "is negative")
// 	}
// }

// func main() {
// Check(1)
// Check(0)
// Check(-1)
// }

// func Positive(n int) (bool, error) {
//     if n == 0 {
//         return false, errors.New("undefined")
//     }
//     return n > -1, nil
// }

// func Check(n int) {
//     pos, err := Positive(n)
//     if err != nil {
//         fmt.Println(n, err)
//         return
//     }
//     if pos {
//         fmt.Println(n, "is positive")
//     } else {
//         fmt.Println(n, "is negative")
//     }
// }
