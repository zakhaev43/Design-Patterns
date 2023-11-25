package main

import (
	"fmt"

	"github.com/zakhaev43/Golang-Practice/Basics"
)

func main() {

  
   var c chan string = make(chan string)

   go Basics.Printer(c)
   go Basics.Pinger(c)
 
   var input string
   fmt.Scanln(&input)





   // Basics.Pstring()

     //Basics.Lc()
    // Basics.ArraysP();
   
    /*
    p:=15
    fmt.Println(p)
    Basics.Ppoint(&p)
    */

    /*
   j:= Basics.Pstr()

   fmt.Println(j)

   */





   }