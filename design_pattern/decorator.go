package design_pattern

import "fmt"

type Coffee interface {
    Cost() int
}

type Espresso struct{}

func (e *Espresso) Cost() int {
    return 10
}

type MilkDecorator struct {
    Coffee Coffee
}

func (m *MilkDecorator) Cost() int {
    return m.Coffee.Cost() + 2
}

func Go_Decorator() {
    espresso := &Espresso{}
    milkAdded := &MilkDecorator{Coffee: espresso}
    fmt.Println("Cost:", milkAdded.Cost())
}
