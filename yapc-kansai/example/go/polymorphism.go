package main

import "fmt"

type Animal interface {
	Say()
}
type Dog string
type Cat string

func (Dog) Say()        { fmt.Println("Waon!!") }
func (Cat) Say()        { fmt.Println("Nyan!!") }
func Say(animal Animal) { animal.Say() }

func main() {
	var dog Dog
	var cat Cat
	Say(dog)
	Say(cat)
}
