package main

import (
	"fmt"
	person "hiper/person/models"
)

func main() {
	//copy -> Person
	//s := person.Init("Oleg", 30, "New York")
	//s.SetAgeValue(10)
	//fmt.Println(s.Age)
	//// pointer -> *0xAddr -> Person
	//i := person.New("John", 30, "New York")
	//i.SetAgeValuePtr(10)
	//fmt.Println(i.Age)

	s := person.InitUser("John", 30)
	s.SetID()
	s.Print()
	fmt.Println(s.Id)

	i := person.NewUser("John", 30)
	i.SetID()
	i.Print()
}
