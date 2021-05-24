package main

import (
	"fmt"
	"strconv"
)

func main()  {

	var a int = 10
	var b int32
	b = int32(a)
	fmt.Println("b=", b)

	var c float32 = 2.7
	var d float64
	d = float64(c)
	fmt.Println("d=", d)

	var e int = int(c)
	fmt.Println("e=", e)

	var f float64 = float64(b)
	fmt.Println("f=", f)

	var g int = 32
	var h string = strconv.Itoa(g)
	fmt.Print("h=")
	fmt.Printf("%q\n", h)

	var i float32 = 10.24
	var j string = fmt.Sprint(i)
	fmt.Print("j=")
	fmt.Printf("%q\n", j)

	var k string ="20"
	l, err := strconv.Atoi(k)
	if err == nil {
		fmt.Println("l=", l)
	}

	var m string = "23.65"
	n, err := strconv.ParseFloat(m, 16)
	fmt.Println("n=", n)

	var o string = "My string"
	var p []byte
	var q string
	p = []byte(o)
	q = string(p)
	fmt.Println("p=", p)
	fmt.Println("q=", q)

}

/*Result:
b= 10
d= 2.700000047683716
e= 2
f= 10
h="32"
j="10.24"
l= 20
n= 23.65
p= [77 121 32 115 116 114 105 110 103]
q= My string
*/