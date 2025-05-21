package person

import "fmt"

type Person struct {
	Name string
	Age  int
}

func New(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func Init(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

func ChageAge(p Person, age int) {
	p.Age = age
}

func ChageAgePtr(p *Person, age int) {
	p.Age = age
}

// Метод
func (p *Person) GetName() string {
	fmt.Println(p.Name)
	return p.Name
}

func (p Person) GetAge() {
	fmt.Println(p.Age)

}
