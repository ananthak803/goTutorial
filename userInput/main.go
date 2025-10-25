package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var name string
	// fmt.Scan(&name)
	reader := bufio.NewReader(os.Stdin)
	name,_:=reader.ReadString('\n')
	fmt.Println("your name is",name);
}
