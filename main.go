package main

import (
	"fmt"
	"hiper/person/models"
)

func main() {

	p := person.SetPerson(18, "Oleg", "Osvita 4")
	fmt.Println(p.Name)
	fmt.Println(p.GetAge())
}
