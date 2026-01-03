package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var operation string
	fmt.Print("Введите операцию (AVG/SUM/MED): ")
	fmt.Scanln(&operation)

	nums := inputUser()
	if len(nums) == 0 {
		fmt.Println("Нет данных для вычислений")
		return
	}

	switch strings.TrimSpace(strings.ToUpper(operation)) {
	case "AVG":
		fmt.Println("Среднее:", avg(nums...))
	case "SUM":
		fmt.Println("Сумма:", sum(nums...))
	case "MED":
		fmt.Println("Медиана:", median(nums...))
	default:
		fmt.Println("Неизвестная операция:", operation)
	}
}

func inputUser() []int {
	fmt.Print("Введите числа через запятую: ")
	var input string
	fmt.Scanln(&input)

	parts := strings.Split(input, ",")
	var nums []int

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		n, err := strconv.Atoi(part)
		if err != nil {
			fmt.Printf("Ошибка: '%s' — не число\n", part)
			return nil
		}
		nums = append(nums, n)
	}

	return nums
}

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func avg(nums ...int) float64 {
	if len(nums) == 0 {
		return 0
	}
	return float64(sum(nums...)) / float64(len(nums))
}

func median(nums ...int) float64 {
	if len(nums) == 0 {
		return 0
	}
	s := make([]int, len(nums))
	copy(s, nums)

	n := len(s)
	if n%2 == 1 {
		return float64(s[n/2])
	}
	return float64(s[n/2-1]+s[n/2]) / 2.0
}
