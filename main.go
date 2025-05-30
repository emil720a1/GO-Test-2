package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	orders := map[string]int{
		"Burger": 2,
		"Pizza":  4,
		"Sushi":  3,
	}

	//n := make([]int, 0, len(orders))
	//for k := range n {
	//	n = append(n, k)
	//}
	ordersReady := make(chan string)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {

		defer wg.Done()
		time.Sleep(time.Duration(orders["Burger"]) * time.Second)
		ordersReady <- "Burger"

	}()

	go func() {

		defer wg.Done()
		time.Sleep(time.Duration(orders["Pizza"]) * time.Second)
		ordersReady <- "Pizza"
	}()

	go func() {
		defer wg.Done()

		time.Sleep(time.Duration(orders["Sushi"]) * time.Second)
		ordersReady <- "Sushi"

	}()

	go func() {
		wg.Wait()
		close(ordersReady)
	}()

	for i := range ordersReady {
		fmt.Println(i)
	}

}
