package main

import "fmt"

func printSeparator() {
	fmt.Println("---------------------")
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

func main() {
	printSeparator()
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
	printSeparator()

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
	printSeparator()
}

/*Result:
---------------------
1+2 = 3
1+2+3 = 6
---------------------
3
7
7
---------------------
*/
