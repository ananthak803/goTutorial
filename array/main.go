package main

import "fmt"

func main() {
	var name [5]string
	name[0] = "anantha"
	name[1] = "krishnan"
	fmt.Println(name) //[anantha krishnan   ]
	fmt.Println(len(name))//5
	fmt.Printf("%q\n",name) //["anantha" "krishnan" "" "" ""]
	var numbers =[3]int{1,2,3}
	fmt.Println(numbers)
	//defualt value for int is 0 and "" for string
	var nums[5]int
	fmt.Println(nums)
}
