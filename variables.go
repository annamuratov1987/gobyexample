package main

import "fmt"

func main()  {

	var a = "initial"
	fmt.Println("a=", a)

	var b, c int = 1, 2
	fmt.Println("b=", b, " c=", c)

	var d = true
	fmt.Println("d=", d)

	var e int
	fmt.Println("e=", e)

	f := "apple"
	fmt.Println("f=", f)

	var g string
	fmt.Println("g=", g)

	var h float32
	fmt.Println("h=", h)

}

/* Result:
a= initial
b= 1  c= 2
d= true
e= 0
f= apple
g=
h= 0
*/