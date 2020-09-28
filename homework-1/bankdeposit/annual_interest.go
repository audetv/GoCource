package bankdeposit

import (
	"fmt"
	"math"
)

func Calculate() {
	var depositAmount float64
	var annualInterest float64
	const period = 5

	fmt.Println("Введите сумму вклада:")
	fmt.Scanln(&depositAmount)

	fmt.Println("Введите годовой процент по вкладу:")
	fmt.Scanln(&annualInterest)

	sum := depositAmount * math.Pow(1+annualInterest/100, period)

	fmt.Printf("Сумма вклада через %x лет: %.2f \n", period, sum)
}
