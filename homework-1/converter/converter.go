package converter

import "fmt"

func Calculate() {
	const dollarRate float64 = 76.26
	var rub float64
	fmt.Println("Введите количество рублей для обмена:")
	fmt.Scanln(&rub)
	sum := rub / dollarRate
	fmt.Print("Вы получите: ")
	fmt.Printf("$%.2f\n\n", sum)
}
