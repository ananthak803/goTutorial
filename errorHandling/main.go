package main

import "fmt"

// func div(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("division by 0") //error string should not start with uppercase 
// 	}
// 	return a/b,nil
// }

func div(a,b int)(int, string){
	if b==0{
		return 0,"Division by zero"
		
	}
	return a/b,"nil"
}

func main() {
	// fmt.Println(div(10,2)) //5 <nil>

	// res,error:=div(10,2) 
	// fmt.Println(res)
	// fmt.Println(error)

	// res,_:=div(10,2)
	// res,_:=div(10,0) //when u need only 1 return value not other use _ for other so no need to use that value.
	// fmt.Println(res)

	res,err:=div(10,0)
	fmt.Println(res)
	fmt.Println(err)
}