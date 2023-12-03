package main

import "fmt"

type Human interface {
	setInfo()
	getInfo()
}

type Person struct {
	name    string
	phone   int
	address string
}

type Employee struct {
	id     int
	dept   string
	salary float64
}

func (e *Employee) setInfo() {

	e.dept = "tax"
	e.id = 12
	e.salary = 5000

}

func (p *Person) setInfo() {
	p.name = "saif"
	p.address = "Bayern"
	p.phone = 123

}

func (e Employee) getInfo() {

	fmt.Println("Id ", e.id, "Phone", e.dept, "Adress", e.salary)

}

func (p Person) getInfo() {

	fmt.Println("Name", p.name, "Address", p.address, "Phone", p.phone)
}

func main() {

	emp := Employee{}
	per := Person{}

	human := []Human{&emp, &per}

	for _, human := range human {
		human.setInfo()
		human.getInfo()

	}
}
