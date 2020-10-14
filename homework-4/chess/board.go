package chess

import (
	"fmt"
	"math"
)

func Process() {

	type Point struct {
		x int
		y int
	}

	var x, y int

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	fmt.Println("Введите позицию x y коня на доске, два цифры от 0 до 7 через пробел.")

	_, err := fmt.Scan(&x, &y)
	if err != nil {
		panic("Неправильные значения позиции")
	}

	current := Point{x, y}
	fmt.Println(current)

	var availablePositions []Point

	for i := x - 2; i <= (x + 2); i++ {
		for j := y - 2; j <= (y + 2); j++ {
			if isPositionValid(x, y, i, j) {
				//selectPosition(i, j)
				//availablePositions
			}
			fmt.Println(isPositionValid(x, y, i, j))
		}
	}

}

func isPositionValid(x int, y int, i int, j int) bool {
	println(x, y, i, j)
	if (math.Abs(float64(x-i)) == 1 && math.Abs(float64(y-j)) == 2) ||
		(math.Abs(float64(x-i)) == 2 && math.Abs(float64(y-j)) == 2) {
		return true
	}
	return false
}
