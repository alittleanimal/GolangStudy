// package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	tty, err := os.OpenFile("test1_pair.go", os.O_RDWR, 0)

	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	var r io.Reader
	r = tty
	fmt.Println(r)

	// var w io.Writer
	// w = r.(io.Writer)

}
