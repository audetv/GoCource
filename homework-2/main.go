package main

import (
	"fmt"
	"strconv"
)

func main() {
	var data string

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	fmt.Println("Введите целое число: ")
	fmt.Scanln(&data)

	num, err := strconv.Atoi(data)

	if err != nil {
		panic(fmt.Sprintf("Вы ввели некорректное число, работа программы остановлена. %v", err))
	}

	doHomework(num)
}

func doHomework(num int) {

	// Задача №1.

	fmt.Println(isEven(num))

	// Задача №2.

	message := "Это число не делится на 3 без остатка"

	if isMultipleOf(3, num) {
		message = "Это число делится на 3 без остатка"
	}
	fmt.Println(message)

	// Задача №3.
	fmt.Println("100 первых чисел Фибоначчи: ")

	fibonacci(100)

}

func isEven(num int) (message string) {
	message = "Вы ввели нечетное число"
	if isMultipleOf(2, num) {
		message = "Вы ввели четное число"
	}
	return
}

// isMultipleOf определяет, делится ли число num без остатка на value
func isMultipleOf(value int, num int) bool {
	if num%value == 0 {
		return true
	}
	return false
}

func fibonacci(i int) {
	var a int64 = 0
	var b int64 = 1
	for n := 0; n < i; n++ {
		a, b = b, a+b
		if n == 0 {
			a = 0
		}
		fmt.Printf("Фибоначчи %v = %v\n", n, a)
	}
}
