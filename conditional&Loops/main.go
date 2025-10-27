package main

import "fmt"

func main() {
	// x := 20
	// //no need for brackets. user bracket for condition like these : x>5  &&(y>5||z<43)
	// if x < 50 {
	// 	fmt.Println("less that 50")
	// }else{
	// 	fmt.Println("greateer")
	// }
	// //similarly else if
	  
	// day:=35
	// switch day{
	// case 1:
	// 	fmt.Println("monday")
	// case 2:
	// 	fmt.Println("tue")
	// case 3:
	// 	fmt.Println("wed")
	// default:
	// 	fmt.Println("other")
	// }

	// //swithc with multiple values in single case
	// month:="jan"
	// switch month{
	// case "feb","jan":
	// 		fmt.Println("yes")
	// case "march","april","may":
	// 	fmt.Println("jjjjjjj")
	// //works without default also
	// }

	// //can add conditions in switch case

	// temp:=30
	// switch {
	// case temp>20:
	// 	fmt.Println("hot")
	// }

	for i:=0;i<5;i++{
		fmt.Println(i)
	}

	//there is only for loop in go no whlie loops
	//how to while loop with for loop
	count:=0
	for{
		fmt.Println("while loop")
		count++;
		if count==5{
			break
		}
	}

	numbers:=[]int{6,3,6,9,2}
	for index,value:=range numbers{
		fmt.Println(index," ",value)
	}
	  
	data:="hello"
	for index,value:=range data{
		fmt.Println(index," ",value)
	}
}