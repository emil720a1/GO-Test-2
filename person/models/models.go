package person

type Person struct {
	Name string
	Age  int
}

func GetPerson() Person {
	return Person{}
}

func (p Person) GetAge() int {
	return p.Age
}

func SetPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}
