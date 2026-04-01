package main

import "fmt"

const (
	usdToEur = 0.94
	usdToRub = 92.50
	eurToRub = usdToRub / usdToEur
)

func main() {

fmt.Println("Курсы валют:")
fmt.Println("USD -> EUR =", usdToEur)
fmt.Println("USD -> RUB =", usdToRub)
fmt.Println("EUR -> Rub =", eurToRub)

}