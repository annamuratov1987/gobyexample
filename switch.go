package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("It`s the weekend")
	default:
		fmt.Println("It`s a weekdey")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It`s before noon")
	default:
		fmt.Println("It`s after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I`m a bool")
		case int:
			fmt.Println("I`m a int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

/*Result:
Write 2 as two
It`s a weekdey
It`s after noon
I`m a bool
I`m a int
Don't know type string
*/
