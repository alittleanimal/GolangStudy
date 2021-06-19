package main

import (
	"fmt"
	"time"
)

func main() {

	// go func() {
	// 	defer fmt.Println("A.deger")

	// 	func() {
	// 		defer fmt.Println("B.deger")
	// 		runtime.Goexit()
	// 		fmt.Println("B")
	// 	}()

	// 	fmt.Println("A")
	// }()

	go func(a int, b int) bool {
		fmt.Println("a = ", a, "b = ", b)
		return true
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}
}
