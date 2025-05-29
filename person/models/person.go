package main

import "fmt"

func main() {
	slice := []uint8{1, 2, 3, 4}
	for i := range uint8(255) {
		slice = append(slice, i)
	}

	lol := slice[:len(slice):len(slice)]

	for i := range slice {
		slice[i] *= 2
	}

	lol[10] = 155

	fmt.Println("lol[10]", lol[10])
	fmt.Println("slice[10]", slice[10])

	fmt.Println(cap(lol))
	fmt.Println(cap(slice))

	lol = append(lol, 255)
	lol[10] = 255

	fmt.Println("lol[10]", lol[10])
	fmt.Println("slice[10]", slice[10])

	fmt.Println("Cap lol", cap(lol))
	fmt.Println("Cap slice", cap(slice))
}
