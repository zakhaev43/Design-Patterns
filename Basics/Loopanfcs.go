package Basics

import "fmt"

func Lc() {
	i := 1
	for i <= 10 {
		fmt.Println("Salam")
		i = i + 1
		if i%2==0{
			i=i+1
		}
	}
}
