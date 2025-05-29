package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	files := []string{"file1", "file2", "file3"}
	fmt.Println(files)
	var wg sync.WaitGroup
	wg.Add(len(files))
	for _, file := range files {
		go func(file string) {
			defer wg.Done()
			fmt.Println("Починаємо завантаження...", file)
			time.Sleep(1 * time.Second)
			fmt.Println("Завершенно...", file)
		}(file)

	}
	wg.Wait()

}
