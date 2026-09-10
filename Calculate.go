package main

import (
	"fmt"
)

func main() {

	var total float64
	var amount float64

	for {
		fmt.Print("Сколько вы потратили?")
		fmt.Scan(&amount)

		if amount == 0 {

			break
		}

		if amount < 0 {

			fmt.Print("Число не может быть отрицательным,введите корректное число.")

			continue
		}

		total += amount
	}

	fmt.Printf("Всего вы потратили: %.2f ", total)

}
