package main

import (
	"fmt"
	"sync"
	"time"
)

func workerWG(id int, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}
func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go workerWG(i, &wg)
	}

	wg.Wait()
}
