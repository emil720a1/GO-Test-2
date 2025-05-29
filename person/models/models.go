package person

import "fmt"

type IPerson interface {
	SetAgeValue(age int)
	SetName(name string)
	GetName() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) SetAgeValue(age int) {
	fmt.Println("SetAgeValue	", age)
	p.Age = age
}

func (p *Person) SetName(name string) {
	fmt.Println("SetName: ", name)
	p.Name = name
}

func (p Person) GetName() string {
	return p.Name
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

type User struct {
	Name string
	Age  int
}

func (u *User) SetAgeValue(age int) {
	u.Age = age
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (p User) GetName() string {
	return p.Name

}

func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}
