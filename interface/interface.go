package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

// --------Rectangle area and perimeter---------
func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

// --------Circle area and perimeter---------
func (c circle) area() float64 {
	return math.Pi*c.radius*c.radius
}
func (c circle) perimeter() float64 {
	return 2*math.Pi*c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(" Area = ", g.area())
	fmt.Println(" Perimeter = ", g.perimeter())
}

func main() {
	r := rectangle{width: 5, height: 10}
	c := circle{radius: 15}

	measure(r)
	measure(c)
}