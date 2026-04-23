package main

import (
	"fmt"
	"strconv"
	"strings"
)

var exchangeRates = map[string]float64{
    "USD": 1.0,
    "EUR": 0.94,
    "RUB": 92.50,
}

func inputCurrency(prompt string) string {
	for {
		fmt.Print(prompt + " (USD, EUR, RUB): ")
		var input string
		fmt.Scan(&input)
		input = strings.ToUpper(input)

		if _, exists := exchangeRates[input]; exists {
			return input
		}
		fmt.Println("Ошибка! Введите правильно: USD, EUR или RUB")
	}
}
func inputAmount() float64 {
	for {
		fmt.Print("Введите сумму: ")
		var input string
		fmt.Scan(&input)

		amount, err := strconv.ParseFloat(input, 64)
		if err == nil && amount >= 0 {
			return amount
		}
		fmt.Println("Ошибка! Введите число (например, 100.50)")
	}
}
func main() {
	fmt.Println("--- Меню конвертации ---")

	
	from := inputCurrency("Выберите исходную валюту")
	
	amount := inputAmount()

	to := inputCurrency("Выберите целевую валюту")

	amountInUSD := amount / exchangeRates[from]
	result := amountInUSD * exchangeRates[to]

	fmt.Printf("\nИтог: %.2f %s = %.2f %s\n", amount, from, result, to)
}
