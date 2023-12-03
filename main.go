package main

import "fmt"

type Person struct{
  name string
  phone int
  location string
}

func (p Person)get_person(){

  fmt.Println(p.name,p.phone,p.location)

}

func (p *Person)set_person(){

  fmt.Scan(&p.name)
  fmt.Scan(&p.phone)
  fmt.Scan(&p.location)
}

func main(){

  per := Person{ "Saif", 123, "De"}

  
  per.get_person()
  per.set_person()
  per.get_person()

}