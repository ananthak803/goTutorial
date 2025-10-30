package main

import "fmt"

func main() {
	fmt.Println("hello")
	defer fmt.Println("world")
	fmt.Println("!!!!!!!!!!")
	// hello
	// !!!!!!!!!!
	// world

	fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three") //multiple defer means function goes to stack and follow lifo
	fmt.Println("four")

	// one
	// four
	// three
	// two
	// world
}