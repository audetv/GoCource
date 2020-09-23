package main

import (
	"fmt"
	"math"
)

func main() {
	a := 1.0
	b := 3.0

	area := a * b / 2
	fmt.Println("Площадь треугольника: ", area)

	hypotenuse := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	perimeter := a + b + hypotenuse

	fmt.Println("Периметр треугольника: ", perimeter)
	fmt.Println("Гипотенуза треугольника: ", hypotenuse)
}
