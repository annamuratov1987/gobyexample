package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main()  {
	fmt.Println("s=", s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println("d=", d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

/*Result:
s= constant
d= 6e+11
600000000000
-0.28470407323754404
*/