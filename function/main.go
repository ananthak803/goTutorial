package main

import "fmt"

//first parenthesis function parameter then return type if only 1 return type no parenthesis needed.
func add(a int, b int) int {
	return a + b
}

//if multiple parameter of same type we can do this add type at end after all params
//in return we can wriet variable name to return with type, and in funciton we calc that variable an at last just write return it returns the variable, your choice to write return or return result both work.
func sum(a,b,c int)(result int){
	result=a+b+c
	return 
}

func main() {
	res := add(3, 2)
	fmt.Println(res)
	fmt.Println(sum(1,1,1))
}