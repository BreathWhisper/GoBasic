package main

import (
	"errors"
	"fmt"
	"strings"
)

var CONVERT = map[string]map[string]float64{
	"EUR": {
		"USD": 1.1719,
		"RUB": 91.6579,
		"EUR": 1,
	},
	"USD": {
		"EUR": 0.8533,
		"RUB": 78.2117,
		"USD": 1,
	},
	"RUB": {
		"USD": 0.0128,
		"EUR": 0.0109,
		"RUB": 1,
	},
}

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

	result := currencyAmount * CONVERT[fromCurrency][toCurrency]

	return result
}

func checkCurrencyInput() (string, error) {
	var currency string
	fmt.Scan(&currency)
	currency = strings.ToUpper(currency)
	if currency == "USD" || currency == "EUR" || currency == "RUB" {
		return currency, nil
	}
	return "", errors.New("неверный ввод, попробуйте USD, EUR, RUB")
}

func checkCurrencyAmmountInput() (float64, error) {
	var ammount float64
	fmt.Scan(&ammount)
	if ammount > 0 {
		return ammount, nil
	}
	return 0, errors.New("количество денег не может быть отрицательным или строкой")
}

func checkRepeatCalculation() bool {
	var answer string

	fmt.Println("Хотите провести еще один рассчет (Y/N)")
	fmt.Scan(&answer)

	return strings.ToUpper(answer) == "Y"
}

func outputResult(r float64) {
	fmt.Printf("Результат конвертации будет равен = %.3f \n", r)
}
