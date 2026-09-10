package main

import (
	"fmt"
)

func main() {

	var total float64
	var amount float64
	var theNumberOfPurchases = 0
	var averageCheck float64 = 0

	fmt.Println("Калькулятор расходов (вводите числа, чтобы завершить работу введите - 0).")

	for {
		fmt.Print("Сколько вы потратили?")
		fmt.Scan(&amount)

		if amount == 0 {

			break
		}

		if amount < 0 {

			fmt.Println("Число не может быть отрицательным,введите корректное число.")

			continue
		}

		theNumberOfPurchases++

		total += amount

	}

	fmt.Printf("\nВсего вы потратили: %.2f.\n", total)
	fmt.Printf("Всего транзакций совершено: %d.\n", theNumberOfPurchases)

	if theNumberOfPurchases > 0 {

		averageCheck = total / float64(theNumberOfPurchases)
		fmt.Printf("Ваш средний чек: %.2f.", averageCheck)
	} else {
		fmt.Println("У вас не было трат.")
	}
}
