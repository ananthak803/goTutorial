package main

import "fmt"

func byrefernce(num * int){
	*num=*num+10
}

func main() {
	var num int=10

	// num:=10
	// ptr:=&num
	var ptr *int = &num
	fmt.Println(num)
	fmt.Println(ptr)
	fmt.Println(*ptr)

	var p1* int
	fmt.Println(p1)//<nil>
	// fmt.Println(*p1) //this line will throw error
	
	// byrefernce(&num)
	byrefernce(ptr)
	fmt.Println(num)
}
