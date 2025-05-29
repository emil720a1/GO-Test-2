package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan string, 3)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		ch <- "McOleg"
	}()

	go func() {
		defer wg.Done()
		ch <- "Zina"
	}()

	go func() {
		defer wg.Done()
		ch <- "love"
	}()
	go func() {
		wg.Wait()
		close(ch)
	}()

	for name := range ch {
		fmt.Println(name)
	}

	//ch1 := make(chan string, 3)
	//
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//
	//	ch1 <- "oleg"
	//	//fmt.Println(fmt.Sprintf("Привіт %s", i))
	//}()
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//
	//	ch1 <- "zina"
	//
	//}()
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//
	//	ch1 <- "Koncha"
	//}()
	//fmt.Println(fmt.Sprintf("Hello %s", <-ch1))
	//fmt.Println(fmt.Sprintf("Hello %s", <-ch1))
	//fmt.Println(fmt.Sprintf("Hello %s", <-ch1))
	//
	//wg.Wait()

}
