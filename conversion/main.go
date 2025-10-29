package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 20
	fmt.Printf("%T\n", num)
	//int to float
	var data float64 = float64(num)
	fmt.Printf("%.3f\n", data)
	fmt.Printf("%T\n", data)

	//number to string
	str := strconv.Itoa(num)
	fmt.Printf("%s\n", str)
	fmt.Printf("%T\n", str)

	//string to number
	stringNum:="1234"
	numInt,_:=strconv.Atoi(stringNum)
	fmt.Printf("%d\n", numInt)
	fmt.Printf("%T\n", numInt)

	//string to float
	stringFloat:="3.14"
	//only work when string can be converted to integer only
	// numFloat,_:=strconv.Atoi(stringFloat)
	// fmt.Println(numFloat)//0
	// fmt.Printf("%T\n", numInt)

	numFloat,_:=strconv.ParseFloat(stringFloat,64)
	fmt.Println(numFloat)
	fmt.Printf("%T\n", numFloat)


}
