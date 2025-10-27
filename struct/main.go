package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Contact struct{
	email string
	phone string
}

//complex struct
type Emp struct{
	emp_details Person
	emp_contact Contact
}

func main() {
	// var p1 Person // or we can use new keyword "var p1=new(Person)"
	var p1=new(Person)
	p1.name = "anantha"
	p1.age = 12
	fmt.Println(p1) //{anantha 12}

	p2:=Person{
		name:"goat",
		age:15,
	}
	fmt.Println(p2)//{goat 15}

	var emp1 Emp
	emp1.emp_details=Person{
		name:"abc",
		age:20,
	}
	emp1.emp_contact=Contact{
		phone:"12343455",
		email:"123@gmail.com",
	}

	fmt.Println(emp1)//{{abc 20} {123@gmail.com 12343455}}

}
