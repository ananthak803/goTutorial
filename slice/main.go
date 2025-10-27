package main

import "fmt"

func main() {
	//creating is same as array but size is omitted
	// num := []int{} //empty slice
	//make takes 3 arguemtn, type, length and capacity
	num:=make([]int,0)
	// var num=make([]int,0)
	fmt.Println(num)
	fmt.Println(len(num))
	fmt.Println(cap(num))
	num = append(num, 1, 2, 3, 4)
	fmt.Println(num)
	fmt.Println(len(num))
	fmt.Println(cap(num))
}