package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() { messages <- "ping 1" }()
	go func() { messages <- "ping 2" }()

	msg := <-messages
	fmt.Println(msg)
}

/*Result
ping 2
*/
