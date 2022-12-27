package main

import (
	"fmt"
)

type Fish struct {}

func (Fish) Speak() {
	fmt.Println("Glu glu...")
}

type Dog struct {}

func (Dog) Speak() {
	fmt.Println("Woof woof...")
}

type Animal interface {
	Speak() //Receives and returns nothing
}

func makeItSpeak(a Animal) {
	a.Speak()
}

func main(){
	fmt.Println("Hello world")

	f := Fish{}
	f.Speak()
	d := Dog{}
	d.Speak()
}