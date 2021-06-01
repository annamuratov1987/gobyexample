package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rectngle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (rect rectngle) area() float64 {
	return rect.width * rect.height
}

func (rect rectngle) perim() float64 {
	return 2*rect.width + 2*rect.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area:", g.area())
	fmt.Println("perim:", g.perim())
}

func main() {
	r := rectngle{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}

/*Result:
{3 4}
area: 12
perim: 14
{5}
area: 78.53981633974483
perim: 31.41592653589793
*/
