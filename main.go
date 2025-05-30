package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sentences := []string{
		"Go is fast.",
		"Goroutines are light.",
		"Channels sync them.",
	}

	ch := make(chan rune)
	var wg sync.WaitGroup

	for _, sentence := range sentences {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			for _, char := range s {
				time.Sleep(1 * time.Second)
				ch <- char
			}
		}(sentence)
	}

	// Закриваємо канал, коли всі горутини завершились
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Виводимо символи по одному
	for char := range ch {
		fmt.Printf("%c\n", char)
	}
}
