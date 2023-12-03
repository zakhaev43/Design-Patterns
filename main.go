package main

import "fmt"

type Person struct{

  name string
  phone string
  address string
}

type Employee struct{
  Person //embedding

  id int
  dept string
  salary float64

}

func (e *Employee)setInfo(){

  fmt.Scan(&e.name)
  fmt.Scan(&e.phone)
  fmt.Scan(&e.address)
  fmt.Scan(&e.salary)


}

func (e Employee)getInfo(){

  fmt.Print(e.id, e.name, e.phone,e.salary )

}


func main(){

    emp:=Employee{}

    emp.setInfo()
    emp.getInfo()
 

}