package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func calculate(operation string, input string) (float64, error) {
	parts := strings.Split(input, ",")
	var numbers []float64
for _, p := range parts {
	num, err := strconv.ParseFloat(strings.TrimSpace(p), 64)
		if err != nil {
			return 0, fmt.Errorf("ошибка в числе '%s': %v", p, err)
		}
		numbers = append(numbers, num)
	}

	if len(numbers) == 0 {
		return 0, fmt.Errorf("список чисел пуст")
	}

	operation = strings.ToUpper(strings.TrimSpace(operation))

	switch operation {
	case "SUM":
		sum := 0.0
		for _, n := range numbers {
			sum += n
		}
		return sum, nil

	case "AVG":
		sum := 0.0
		for _, n := range numbers {
			sum += n
		}
		return sum / float64(len(numbers)), nil

	case "MED":
		sort.Float64s(numbers)
		n := len(numbers)
		if n%2 == 1 {
			return numbers[n/2], nil
		}
		return (numbers[n/2-1] + numbers[n/2]) / 2, nil

	default:
		return 0, fmt.Errorf("неизвестная операция: %s", operation)
	}
}

func main() {
	var op, input string

	fmt.Print("Введите операцию (AVG, SUM, MED): ")
	fmt.Scanln(&op)

	fmt.Print("Введите числа через запятую: ")
	var line string
	fmt.Scanln(&line)
	input = line 

	result, err := calculate(op, input)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Printf("Результат: %.2f\n", result)
}