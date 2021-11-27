package main

import (
	"fmt"
)

type Person struct {
	minAge int
	name   string
	age    int
}

func (p Person) Hello() {
	fmt.Printf("Hello. My name is %s\n", p.name)
}

func (p Person) MyAge() {
	fmt.Printf("My age is %d\n", p.age)
}

func main() {
	js := Person{minAge: 0, name: "yundream", age: 33}
	js.Hello()
	js.MyAge()
}

func New(minAge int, name string, int age) *Person {
	return &Person{minAge: minAge, name: name, age: age}
}

func Open(name string) (file *File, err error)

f, err := os.Open("filename.txt")
if err != nil {
    fmt.Println("File open error : ", err.Error())
    os.Exit(1)
}