package main

import "fmt"

// 方法1，2，3可以声明全局变量

func main() {
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)

	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcd"
	fmt.Printf("cc = %s, type of cc = %T\n", cc, cc)

	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	g := 3.14 // 八个字节浮点型
	fmt.Println("g = ", g)
	fmt.Printf("type of g = %T\n", g)

	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, "yy = ", yy)

	var kk, ll = 100, "abcd"
	fmt.Println("kk = ", kk, "ll = ", ll)

	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, "jj = ", jj)
}
