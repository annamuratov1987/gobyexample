package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))
	fmt.Println("capacity:", cap(s))

	s = append(s, "d")
	fmt.Println("append:", s)
	fmt.Println("len:", len(s))
	fmt.Println("capacity:", cap(s))

	s = append(s, "e", "f")
	fmt.Println("append:", s)
	fmt.Println("len:", len(s))
	fmt.Println("capacity:", cap(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d", twoD)
}

/*Result:
emp: [  ]
set: [a b c]
get: c
len: 3
capacity: 3
append: [a b c d]
len: 4
capacity: 6
append: [a b c d e f]
len: 6
capacity: 6
copy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl [g h i]
2d [[0] [1 2] [2 3 4]]
*/
