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
	firstTriangle := Triangle{10, 20}
	firstTrapeze := Trapeze{15, 20, 10}

	var firstShape Shape = firstTriangle
	var secondShape Shape = firstTrapeze

	fmt.Println(firstShape)
	fmt.Println(secondShape)

	total := SumAreas(firstShape, secondShape)
	fmt.Println(total)

}

func SumAreas(shapes ...Shape) float64 {
	res := 0.0
	for _, shape := range shapes {
		res += shape.Area()
	}
	return res
}
