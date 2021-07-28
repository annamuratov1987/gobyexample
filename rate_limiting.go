package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("-----------------------")

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

/*Result:

request 1 2021-07-28 10:39:46.42704729 +0500 PKT m=+0.200524058
request 2 2021-07-28 10:39:46.627487477 +0500 PKT m=+0.400964305
request 3 2021-07-28 10:39:46.826800214 +0500 PKT m=+0.600277056
request 4 2021-07-28 10:39:47.027348427 +0500 PKT m=+0.800825257
request 5 2021-07-28 10:39:47.226651263 +0500 PKT m=+1.000128103
-----------------------
request 1 2021-07-28 10:39:47.226830799 +0500 PKT m=+1.000307600
request 2 2021-07-28 10:39:47.226862391 +0500 PKT m=+1.000339166
request 3 2021-07-28 10:39:47.226882756 +0500 PKT m=+1.000359539
request 4 2021-07-28 10:39:47.42726236 +0500 PKT m=+1.200739215
request 5 2021-07-28 10:39:47.627749607 +0500 PKT m=+1.401226452

*/
