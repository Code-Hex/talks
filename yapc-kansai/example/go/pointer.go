package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	obj := New()
	obj.MyNameIs() // "codehex"
	obj.CodeHexToPapix()
	obj.MyNameIs() // "papix"
}

func New() *Person {
	return &Person{
		Name: "codehex",
		Age:  21,
	}
}

func (p *Person) MyNameIs() {
	fmt.Println(p.Name)
}

func (p *Person) CodeHexToPapix() {
	p.Name = "papix"
}
