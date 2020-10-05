package main

import (
	"fmt"
	"strconv"
)

func main() {

	var data string
	var number = 0

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	for {
		fmt.Println("Введите целое число, или наберите «exit» для выхода: ")
		fmt.Scanln(&data)

		if data == "exit" {
			panic(nil)
		}

		num, err := strconv.Atoi(data)

		if err == nil {
			number = num
			break
		}
		fmt.Println("Не удалось распознать число, поробуйте еще раз.")
	}

	doHomework(number)

}

func doHomework(num int) {

	fmt.Println("Задача №1.")
	fmt.Println(isEven(num))

	fmt.Println("Задача №2.")

	message := "Это число не делится на 3 без остатка"
	if isMultipleOf(3, num) {
		message = "Это число делится на 3 без остатка"
	}
	fmt.Println(message)

	fmt.Println("Задача №3.")
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
