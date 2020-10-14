package chess

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

var x, y int

var availablePositions []Point

func Process() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	fmt.Println("Введите позицию x y коня на доске, два цифры от 0 до 7 через пробел.")

	_, err := fmt.Scan(&x, &y)
	if err != nil {
		panic("Неправильное значение позиции")
	}

	// Проходим по всем позициям на доске x,y +2/-2 от введенной позиции.
	for i := x - 2; i <= (x + 2); i++ {
		for j := y - 2; j <= (y + 2); j++ {
			if isPositionValid(i, j) {
				makePositionAvailable(i, j)
			}
		}
	}

	fmt.Println(availablePositions)

}

func makePositionAvailable(i int, j int) {
	availablePositions = append(availablePositions, Point{i, j})
}

// Todo сложное условие, подумать как упростить запись.
func isPositionValid(i int, j int) bool {

	if ((math.Abs(float64(x-i)) == 1 && math.Abs(float64(y-j)) == 2) ||
		(math.Abs(float64(x-i)) == 2 && math.Abs(float64(y-j)) == 1)) &&
		(i > -1 && j > -1 && i < 8 && j < 8) {
		return true
	}
	return false
}
