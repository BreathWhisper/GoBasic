package main

import "fmt"

const USDRUB = 78.2117
const USDEUR = 0.8533

func main() {
	const EURRUB = USDRUB / USDEUR

	var currencyAmount float64
	var fromCurrency string
	var toCurrency string

	fmt.Println("Калькулятор валют")
	currencyAmount, fromCurrency, toCurrency = getUserInput()
	converter(currencyAmount, fromCurrency, toCurrency)

}

func getUserInput() (float64, string, string) {
	var currencyAmount float64
	var fromCurrency string
	var toCurrency string
	fmt.Print("Введите количество денег: ")
	fmt.Scan(&currencyAmount)
	fmt.Print("Введите исходную валюту: ")
	fmt.Scan(&fromCurrency)
	fmt.Print("Введите валюту в кторую хотите перевести: ")
	fmt.Scan(&toCurrency)
	return currencyAmount, fromCurrency, toCurrency
}

func converter(currencyAmount float64, fromCurrency string, toCurrency string) float64 {
	return 0.0
}
