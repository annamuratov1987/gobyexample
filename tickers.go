package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

/*Result:
Tick at 2021-06-24 16:07:45.994288451 +0500 PKT m=+0.500884789
Tick at 2021-06-24 16:07:46.493896247 +0500 PKT m=+1.000492582
Tick at 2021-06-24 16:07:46.993676663 +0500 PKT m=+1.500273025
Ticker stopped
*/
