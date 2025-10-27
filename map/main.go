package main

import "fmt"

func main() {
	
	studentGrades := make(map[string]int)
	studentGrades["anantha"] = 100
	studentGrades["krishnan"] = 60
	studentGrades["vishnu"] = 10
	fmt.Println(studentGrades)//map[anantha:100 krishnan:60 vishnu:10]
	fmt.Println(studentGrades["bob"])//0
	//delete from map
	delete(studentGrades,"krishnan")
	fmt.Println(studentGrades)

	//check if key exists
	_,exist:=studentGrades["anantha"]
	fmt.Println(exist)
	grade,exist1:=studentGrades["bob"]
	fmt.Println(grade ," ",exist1)

	//map using map literal
	person:=map[string]int{
		"alice":10,
		"bob":40,
	}
	fmt.Println(person)//map[alice:10 bob:40]

}
