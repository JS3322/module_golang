/*
    * @Process: complete
    * @Project_Name: module
    * @Package_Name: _test
    * @Made_By: JS
    * @The_creation_time: -
    * @File_Name: functionTest.go
	* @Version : go1.17.2
	* @contents: -
*/

package main

import (
	"fmt"
)

//type assignment after compiler
var testString string

//compiler
type Dog struct {
	name string
}

type Cat struct {
	name string
}

//struct
func NewDog(name string) Dog {
	var d Dog
	d.name = name
	return d
}

//public is available outside the package
func (d Dog) Sounds() {
	fmt.Printf("%s : woof! woof!\n", d.name)
}

func NewCat(name string) *Cat {
	var c Cat
	c.name = name
	return &c
}

//class
func (c *Cat) Sounds() {
	fmt.Printf("%s: meow.\n", c.name)
}

type Animal interface {
	Sounds()
}

func main() {
	d1 := NewDog("Baduk")
	d2 := d1
	d2.name = "Nureong"
	d1.Sounds()
	d2.Sounds()

	c1 := NewCat("Happy")
	c2 := c1
	c2.name = "Merry"
	c1.Sounds()
	c2.Sounds()

}
