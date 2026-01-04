package main

import (
	"fmt"
	"strings"
)

// Курсы валют относительно RUB
var rates = map[string]float64{
	"USD": 90.59,
	"EUR": 77.96,
	"RUB": 1.0, // базовая валюта
}

func main() {
	fmt.Println("Конвертер валют")
	fmt.Print("Поддерживаемые валюты: ")
	for currency := range rates {
		fmt.Print(currency, " ")
	}
	fmt.Println()

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
	fromCurrency := inputCurrency("исходную")
	if fromCurrency == "" {
		return
	}

	amount := inputAmount()

	toCurrency := inputCurrency("целевую")
	if toCurrency == "" {
		return
	}

	result := calculate(amount, fromCurrency, toCurrency)
	fmt.Printf("\n✅ %.2f %s = %.2f %s\n", amount, fromCurrency, result, toCurrency)
}

func inputCurrency(kind string) string {
	for {
		fmt.Printf("Введите %s валюту (%s): ", kind, getSupportedCurrencies())
		var currency string
		fmt.Scanln(&currency)
		currency = strings.ToUpper(strings.TrimSpace(currency))

		if _, exists := rates[currency]; exists {
			return currency
		}
		fmt.Println("❌ Неверная валюта. Попробуйте снова.")
	}
}

func getSupportedCurrencies() string {
	currencies := make([]string, 0, len(rates))
	for c := range rates {
		currencies = append(currencies, c)
	}
	return strings.Join(currencies, "/")
}

func inputAmount() float64 {
	for {
		fmt.Print("Введите сумму: ")
		var amount float64
		_, err := fmt.Scanln(&amount)

		if err != nil || amount < 0 {
			fmt.Println("❌ Неверная сумма. Введите положительное число.")
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

	// Конвертация: from → RUB → to
	rub := amount * rates[fromCurrency]
	return rub / rates[toCurrency]
}
