package Basics

import "fmt"



func ArraysP(){
	p := make([]int, 5)

	i:=1

	for i<5 {
		p[i]=i

		fmt.Println(p[i])

		i++;
	}

	

}