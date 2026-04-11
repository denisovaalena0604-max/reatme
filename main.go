package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	usdToEur = 0.94
	usdToRub = 92.50
	eurToRub = usdToRub / usdToEur
)

func inputCurrency(prompt string) string {
	for {
		fmt.Print(prompt + " (USD, EUR, RUB): ")
		var input string
		fmt.Scan(&input)
		input = strings.ToUpper(input)

		if input == "USD" || input == "EUR" || input == "RUB" {
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

	amount := inputAmount()

	from := inputCurrency("Выберите исходную валюту")

	to := inputCurrency("Выберите целевую валюту")

	var result float64

	var amountInUSD float64
	if from == "USD" {
		amountInUSD = amount
	} else if from == "EUR" {
		amountInUSD = amount / usdToEur
	} else {
		amountInUSD = amount / usdToRub
	}

	if to == "USD" {
		result = amountInUSD
	} else if to == "EUR" {
		result = amountInUSD * usdToEur
	} else {
		result = amountInUSD * usdToRub
	}

	fmt.Printf("\nИтог: %.2f %s = %.2f %s\n", amount, from, result, to)
}
