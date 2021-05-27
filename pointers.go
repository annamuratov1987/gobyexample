package main

import "fmt"

func zeroval(ival int) {
	ival = 0
	fmt.Println("zeroval -> ival:", ival)
}

func zeroptr(iptr *int) {
	*iptr = 0
	fmt.Println("zeroval -> iptr:", iptr)
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}

/*Result:
initial: 1
zeroval -> ival: 0
zeroval: 1
zeroval -> iptr: 0xc00001e0b8
zeroptr: 0
pointer: 0xc00001e0b8
*/
