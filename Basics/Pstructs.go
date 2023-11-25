package Basics

import "fmt"

type Person struct {
    Name string
    Age  int
}
func Pstr() Person{

	var a Person;

	a.Name="Mahdi"
	a.Age=40

	fmt.Println(a)
    return a

}