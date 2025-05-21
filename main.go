package main

import (
	person "hiper/person/models"
)

func main() {

	s := person.Init("Oleg", 18)

	person.ChageAge(s, 43)
	s.GetAge()

	i := person.New("John", 41)

	person.ChageAgePtr(i, 43)
	i.GetAge()
}
