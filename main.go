// go mod init gotutorial : init a go project adds go.mod file
package main
import "fmt"
import "gotutorial/testpkg"

// main is entry point of program
func main(){
	fmt.Println("helllooooooo")
	testpkg.TestFunc()
	var step int =2
	testpkg.Step(step)

	var left int =1
	var right int=2
	fmt.Println(left+right)

	var fname string="anantha"
	var lname string="krishnan"
	fmt.Println(fname+" "+lname)

	var is bool =false
	fmt.Println(is)

}

//https://youtu.be/akosxcqJorU?si=cbs0s_5xk7dTgf8i
//1:33:20