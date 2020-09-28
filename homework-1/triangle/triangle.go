package triangle

import (
	"fmt"
	"math"
)

func Calculate() {
	a := 1.0
	b := 3.0

	area := a * b / 2
	fmt.Printf("Стороны треугольника a: %.2f, b: %.2f \n", a, b)

	fmt.Println("Площадь треугольника: ", area)

	hypotenuse := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	perimeter := a + b + hypotenuse

	fmt.Printf("Периметр треугольника: %.2f\n", perimeter)
	fmt.Printf("Гипотенуза треугольника: %.2f\n", hypotenuse)
}
