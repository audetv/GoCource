package main

import "fmt"

func main() {
	const dollarRate float64 = 76.26
	var rub float64
	fmt.Println("Введите количество рублей для обмена:")
	fmt.Scanln(&rub)
	sum := rub / dollarRate
	fmt.Printf("%.2f", sum)
}
