package testpkg

import "fmt"

// not exported because function starting letter is lowercase but can be used inside this package
func testfunc(){
	fmt.Println("test function no export")
}

//to make a function export make the first letter in function name capital
func TestFunc(){
	fmt.Println("exported function")
}

// func Step(step int){
// 	if (step==1){
// 		fmt.Println("step 1")
// 	}else if step ==2{            //else if in next line give error
// 		fmt.Println("step2")
// 	}
// }

func Step(step int){
	fmt.Println("step ",step)
}