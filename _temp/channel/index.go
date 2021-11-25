package channel

// import (
// 	"fmt"
// 	"time"
// )

// func n(ch chan int, quit chan bool, s int, e int) {
// 	for i := s; i < e; i++ {
// 		time.Sleep(time.Second)
// 		ch <- i
// 	}
// 	quit <- true
// }

// func main() {
// 	ch, quit := make(chan int), make(chan bool)

// 	go n(ch, quit, 0, 10)

// 	for {
// 		select {
// 		case x := <-ch:
// 			fmt.Println(x)
// 		case <-quit:
// 			return
// 		}
// 	}
// }

// func main() {
// 	// 정수형 채널을 생성한다
// 	ch := make(chan int)

// 	go func() {
// 	  ch <- 123   //채널에 123을 보낸다
// 	}()

// 	var i int
// 	i = <- ch  // 채널로부터 123을 받는다
// 	println(i)
//   }

// import "fmt"

// func main() {
//     done := make(chan bool)
//     go func() {
//         for i := 0; i < 10; i++ {
//             fmt.Println(i)
//         }
//         done <- true
//     }()

//     // 위의 Go루틴이 끝날 때까지 대기
//     <-done
// }

//channel buffer
// // import "fmt"

// // func main() {
// //   c := make(chan int)
// //   c <- 1   //수신루틴이 없으므로 데드락
// //   fmt.Println(<-c) //코멘트해도 데드락 (별도의 Go루틴없기 때문)

// ch := make(chan int, 1)

// //수신자가 없더라도 보낼 수 있다.
// ch <- 101

// fmt.Println(<-ch)
// // }

//channel parameter
// import "fmt"

// func main() {
//     ch := make(chan string, 1)
//     sendChan(ch)
//     receiveChan(ch)
// }

// func sendChan(ch chan<- string) {
//     ch <- "Data"
//     // x := <-ch // 에러발생
// }

// func receiveChan(ch <-chan string) {
//     data := <-ch
//     fmt.Println(data)
// }

//chanel close
// func main() {
//     ch := make(chan int, 2)

//     // 채널에 송신
//     ch <- 1
//     ch <- 2

//     // 채널을 닫는다
//     close(ch)

//     // 채널 수신
//     println(<-ch)
//     println(<-ch)

//     if _, success := <-ch; !success {
//         println("더이상 데이타 없음.")
//     }
// }

// func main() {
//     ch := make(chan int, 2)

//     // 채널에 송신
//     ch <- 1
//     ch <- 2

//     // 채널을 닫는다
//     close(ch)

//     // 방법1
//     // 채널이 닫힌 것을 감지할 때까지 계속 수신
//     /*
//     for {
//         if i, success := <-ch; success {
//             println(i)
//         } else {
//             break
//         }
//     }
//     */

//     // 방법2
//     // 위 표현과 동일한 채널 range 문
//     for i := range ch {
//         println(i)
//     }
// }

//channel select
// import "time"

// func main() {
//     done1 := make(chan bool)
//     done2 := make(chan bool)

//     go run1(done1)
//     go run2(done2)

// EXIT:
//     for {
//         select {
//         case <-done1:
//             println("run1 완료")

//         case <-done2:
//             println("run2 완료")
//             break EXIT
//         }
//     }
// }

// func run1(done chan bool) {
//     time.Sleep(1 * time.Second)
//     done <- true
// }

// func run2(done chan bool) {
//     time.Sleep(2 * time.Second)
//     done <- true
// }
