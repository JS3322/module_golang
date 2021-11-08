package main

import "fmt"

//duck type test
type Animal interface {
	Sounds()
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func NewDog(name string) Dog {
	var d Dog
	d.name = name
	return d
} // 구조체

func (d Dog) Sounds() {
	fmt.Printf("%s: woof woof!\n", d.name)

}

func NewCat(name string) *Cat {
	var c Cat
	c.name = name
	return &c
}

func (c *Cat) Sounds() {
	fmt.Printf("%s: meow.\n", c.name)
} //클래스

func SayHello(a Animal) {
	println("Hello!")
	a.Sounds()
}

func main() {
	var a Animal

	a = NewDog("Baduk")
	SayHello(a)

	a = NewCat("Happy")
	SayHello(a)

	SayHello(NewDog("Nureong"))
}
