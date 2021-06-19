package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called....")
	fmt.Println(arg)

	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
		fmt.Printf("value type is %T, value is %v \n", arg, arg)
	} else {
		fmt.Println("arg is string type, value = ", value)
	}

	fmt.Println("--------")
}

type Book struct {
	auth string
}

func main() {
	book := Book{"Golang"}

	myFunc(book)
	myFunc(100)
	myFunc("abc")
}
