package main

import "fmt"

const (
	usdToEur = 0.94
	usdToRub = 92.50
	eurToRub = usdToRub / usdToEur
)

func getUserInput(promt string) string {
	var input string
	fmt.Print(promt)
	fmt.Scan(&input)
	return input
}

func calculateConversion(amount string, fromCurrency string, toCurrency string) {
	fmt.Printf("Расчет: %s из %s в %s (функция пока пуста)\n", amount, fromCurrency, toCurrency)
}
func main() {

	fmt.Println("Курсы валют:")
	fmt.Println("USD -> EUR =", usdToEur)
	fmt.Println("USD -> RUB =", usdToRub)
	fmt.Println("EUR -> Rub =", eurToRub)
	fmt.Println("-------------------------")
	amount := getUserInput("Введите сумму для конвертации: ")
	fromVal := getUserInput("Введите исходную валюту (например, USD): ")
	toVal := getUserInput("Введите целевую валюту (например, EUR): ")
	calculateConversion(amount, fromVal, toVal)
}
