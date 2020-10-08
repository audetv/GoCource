package main

import (
	"fmt"
)

type Triangle struct {
	a float64
	h float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.a * t.h
}

type Trapeze struct {
	a float64
	b float64
	h float64
}

type Shape interface {
	Area() float64
}

func main() {
	firstTriangle := Triangle{10, 20}
	firstTrapeze := Trapeze{15, 20, 10}
	firstSquare := Square{10}

	total := SumAreas(firstCircle, firstSquare, secondCircle)

	fmt.Println(total)

}

func SumAreas(shapes ...Shape) float64 {
	res := 0.0
	for _, shape := range shapes {
		res += shape.Area()
	}
	return res
}
