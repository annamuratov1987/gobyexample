package main

import "fmt"

func main() {
	m := map[string]int{}
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	value, prs := m["k2"]
	fmt.Println("prs:", prs)
	fmt.Println("value:", value)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

/*Result:
map: map[k1:7 k2:13]
v1: 7
len: 2
map: map[k1:7]
prs: false
value: 0
map: map[bar:2 foo:1]
*/
