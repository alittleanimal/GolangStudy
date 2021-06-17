package main

import "fmt"

/*
func swap(a int, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}
*/

func swap(pa *int, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}

func main() {
	var a int = 10
	var b int = 20

	swap(&a, &b)

	fmt.Println("a = ", a, " b = ", b)

	var p *int
	p = &a
	fmt.Println(p)

	var pp **int = &p
	fmt.Println(pp)
}
