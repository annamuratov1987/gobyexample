package main

import "os"

func main() {

	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

/*Result:

panic: a problem

goroutine 1 [running]:
main.main()
        /home/javlon/projects/gobyexample/panic.go:7 +0x39

Process finished with exit code 2

*/
