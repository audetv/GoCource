package shape

import "fmt"

type Shape interface {
	Area() float64
}

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

func (t Trapeze) Area() float64 {
	return 0.5 * t.h * (t.a + t.b)
}

func Process() {
	var firstTriangle Shape = Triangle{10, 20}
	// Я так понимаю, что могу явно не объявлять, не указвать что firstTrapeze - это Shape.
	firstTrapeze := Trapeze{15, 20, 10}

	// var firstShape Shape = firstTriangle.
	// var secondShape Shape = firstTrapeze.
	// total := SumAreas(firstShape, secondShape),
	// т.е. это более подробно и явно видно, что передаем Shape.

	total := SumAreas(firstTriangle, firstTrapeze)
	fmt.Println(total)

}

func SumAreas(shapes ...Shape) float64 {
	res := 0.0
	for _, shape := range shapes {
		res += shape.Area()
	}
	return res
}
