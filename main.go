package main

import (
	"fmt"
)

// Функция для получения правильного числа от пользователя
func getAmount() float64 {

	var amount float64
	fmt.Print("Сколько вы потратили?")
	fmt.Scan(&amount)
	return amount
}

// Функция для расчёта среднего чека
func calculateAverage(total float64, theNumberOfPurchases int) float64 {
	if theNumberOfPurchases == 0 {
		return 0
	}

	return total / float64(theNumberOfPurchases)
}

func main() {

	var total float64
	var theNumberOfPurchases int

	fmt.Println("Калькулятор расходов (вводите числа, чтобы завершить работу введите - 0).")

	for {

		amount := getAmount()

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

		average := calculateAverage(total, theNumberOfPurchases)
		fmt.Printf("Ваш средний чек: %.2f.", average)
	} else {
		fmt.Println("У вас не было трат.")
	}
}
