package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "anantha,krishnan,vishnu,dragon"
	names := strings.Split(data, ".")
	fmt.Println(names)//[anantha,krishnan,vishnu,dragon]

	str:="one two three one two two five"
	count:=strings.Count(str,"two")
	fmt.Println(count)//3

	s:="       hello  world          "
	fmt.Println(s)
	trimmed:=strings.TrimSpace(s)//remove space from ends but not from middle of string
	fmt.Println(trimmed)//hello  world

	s1:="anantha"
	s2:="krishnan"
	res:=strings.Join([]string{s1,s2},"@")
	fmt.Println(res)//anantha@krishnan
	res2:=strings.Join([]string{s1,s2,"goat"},"---")
	fmt.Println(res2)//anantha---krishnan---goat
}
