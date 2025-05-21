package person

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address string
}

func New(name string, age int, address string) *Person {
	return &Person{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

func Init(name string, age int, address string) Person {
	return Person{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

func (p Person) SetAgeValue(age int) {
	p.Age = age
	fmt.Println(p.GetAge())
}

func (p *Person) SetAgeValuePtr(age int) {
	p.Age = age
	fmt.Println(p.Age)
}

func (p Person) GetAge() int {
	return p.Age
}
