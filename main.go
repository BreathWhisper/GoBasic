package main

import (
	"errors"
	"fmt"
	"strings"
)

const USDRUB = 78.2117
const USDEUR = 0.8533
const EURRUB = USDRUB / USDEUR

func main() {
	fmt.Println("Калькулятор валют")

	var currencyAmount float64
	var fromCurrency string
	var toCurrency string

	var result float64

	for {

		currencyAmount, fromCurrency, toCurrency = getUserInput()
		result = converter(currencyAmount, fromCurrency, toCurrency)
		outputResult(result)

		isRepeatCalculation := checkRepeatCalculation()
		if !isRepeatCalculation {
			break
		}
	}

}

func getUserInput() (float64, string, string) {
	var currencyAmount float64
	var fromCurrency string
	var toCurrency string
	var err error

	for {
		fmt.Print("Введите исходную валюту: ")
		fromCurrency, err = checkCurrencyInput()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Print("Введите количество денег: ")
		currencyAmount, err = checkCurrencyAmmountInput()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Print("Введите валюту в кторую хотите перевести: ")
		toCurrency, err = checkCurrencyInput()
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}
	return currencyAmount, fromCurrency, toCurrency
}

func converter(currencyAmount float64, fromCurrency string, toCurrency string) float64 {
	var result float64

	switch {
	case fromCurrency == "USD" && toCurrency == "EUR":
		result = currencyAmount * USDEUR
	case fromCurrency == "EUR" && toCurrency == "USD":
		result = currencyAmount / USDEUR
	case fromCurrency == "RUB" && toCurrency == "USD":
		result = currencyAmount / USDRUB
	case fromCurrency == "RUB" && toCurrency == "EUR":
		result = currencyAmount / EURRUB
	case fromCurrency == "USD" && toCurrency == "RUB":
		result = currencyAmount * USDRUB
	case fromCurrency == "EUR" && toCurrency == "RUB":
		result = currencyAmount * EURRUB
	default:
		result = currencyAmount
	}

	return result
}

func checkCurrencyInput() (string, error) {
	var currency string
	fmt.Scan(&currency)
	currency = strings.ToUpper(currency)
	if currency == "USD" || currency == "EUR" || currency == "RUB" {
		return currency, nil
	}
	return "", errors.New("Неверный ввод, попробуйте USD, EUR, RUB")
}

func checkCurrencyAmmountInput() (float64, error) {
	var ammount float64
	fmt.Scan(&ammount)
	if ammount > 0 {
		return ammount, nil
	}
	return 0, errors.New("Количество денег не может быть отрицательным или строкой")
}

func checkRepeatCalculation() bool {
	var answer string

	fmt.Println("Хотите провести еще один рассчет (Y/N)")
	fmt.Scan(&answer)
	if strings.ToUpper(answer) == "Y" {
		return true
	}
	return false
}

func outputResult(r float64) {
	fmt.Printf("Результат конвертации будет равен = %.3f \n", r)
}
