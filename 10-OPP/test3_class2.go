// package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human

	level int
}

func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) Print() {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
	fmt.Println("level = ", this.level)
}

func main() {
	h := Human{"zhang3", "female"}

	h.Eat()
	h.Walk()

	fmt.Println("----------")

	// s := SuperMan{Human{"li4", "female"}, 88}
	var s SuperMan
	s.name = "li4"
	s.sex = "male"
	s.level = 88

	s.Walk()
	s.Eat()
	s.Fly()

	s.Print()
}
