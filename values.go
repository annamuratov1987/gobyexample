package main

import "fmt"

func main()  {
	fmt.Println("go"+"lang")
	fmt.Println(`go`+`lang`)
	fmt.Println('g')
	fmt.Println('o')
	fmt.Println('g'+'o')

	fmt.Println("1+1=", 1+1)
	fmt.Println("11/3=", 11/3)
	fmt.Println("11.0/3.0=", 11.0/3.0)

	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)
}

/* Result:
golang
golang
103
111
214
1+1= 2
11/3= 3
11.0/3.0= 3.6666666666666665
true
true
false
 */