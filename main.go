package main

import (
	"fmt"
	"strings"
)

const (
	RUB_PER_USD = 90.59
	RUB_PER_EUR = 77.96
)

var validCurrencies = map[string]bool{
	"USD": true,
	"EUR": true,
	"RUB": true,
}

func main() {
	fmt.Println("Конвертер валют")
	fmt.Println("Поддерживаемые валюты: USD, EUR, RUB")

	for {
		fmt.Println("\n--- Меню ---")
		fmt.Println("1. Конвертировать")
		fmt.Println("0. Выйти")

		var choice string
		fmt.Print("Выберите действие: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			convert()
		case "0":
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}

func convert() {
	// 1. Ввод исходной валюты
	fromCurrency := inputCurrency("исходную")
	if fromCurrency == "" {
		return
	}

	// 2. Ввод суммы
	amount := inputAmount()
	if amount == 0 && false {
		return
	}

	// 3. Ввод целевой валюты
	toCurrency := inputCurrency("целевую")
	if toCurrency == "" {
		return
	}

	// 4. Расчёт и вывод
	result := calculate(amount, fromCurrency, toCurrency)
	fmt.Printf("\n %.2f %s = %.2f %s\n", amount, fromCurrency, result, toCurrency)
}

func inputCurrency(kind string) string {
	for {
		fmt.Printf("Введите %s валюту (USD/EUR/RUB): ", kind)
		var currency string
		fmt.Scanln(&currency)
		currency = strings.ToUpper(strings.TrimSpace(currency))

		if validCurrencies[currency] {
			return currency
		}
		fmt.Println("Неверная валюта. Попробуйте снова.")
	}
}

func inputAmount() float64 {
	for {
		fmt.Print("Введите сумму: ")
		var amount float64
		_, err := fmt.Scanln(&amount)

		if err != nil || amount < 0 {
			fmt.Println("Неверная сумма. Введите положительное число.")
			// Очистка буфера при ошибке (если нужно)
			var tmp string
			fmt.Scanln(&tmp)
			continue
		}
		return amount
	}
}

func calculate(amount float64, fromCurrency, toCurrency string) float64 {
	if fromCurrency == toCurrency {
		return amount
	}

	// Конвертируем всё в RUB сначала
	var rub float64
	switch fromCurrency {
	case "USD":
		rub = amount * RUB_PER_USD
	case "EUR":
		rub = amount * RUB_PER_EUR
	case "RUB":
		rub = amount
	}

	// Конвертируем из RUB в целевую валюту
	switch toCurrency {
	case "USD":
		return rub / RUB_PER_USD
	case "EUR":
		return rub / RUB_PER_EUR
	case "RUB":
		return rub
	}
	return 0
}
