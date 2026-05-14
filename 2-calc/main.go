package main

import (
	"fmt"
	"sort"    // Инструменты для сортировки (нужны для медианы).
	"strconv" // Преобразование строк в числа (string to conversion).
	"strings" // Работа со строками (разбиение, удаление пробелов).
)

func calculate(operation string, input string) (float64, error) {
	parts := strings.Split(input, ",") // Разрезаем строку по запятой. Получаем срез (массив) строк.
	var numbers []float64              // Создаем пустой список, где будем хранить уже готовые числа.

	for _, p := range parts { // Перебираем каждую часть (p), которую получили после разделения.
		// Убираем лишние пробелы и превращаем строку в число типа float64.
		num, err := strconv.ParseFloat(strings.TrimSpace(p), 64)
		if err != nil {
			// Если не удалось превратить в число (например, ввели "abc"), возвращаем ошибку.
			return 0, fmt.Errorf("ошибка в числе '%s': %v", p, err)
		}
		numbers = append(numbers, num) // Добавляем успешно распознанное число в наш список.
	}

	if len(numbers) == 0 { // Проверка на пустоту.
		return 0, fmt.Errorf("список чисел пуст")
	}

	operation = strings.ToUpper(strings.TrimSpace(operation))
	// Мы приводим название операции к верхнему регистру, чтобы "sum" и "SUM" работали одинаково.
	switch operation {
	case "SUM":
		sum := 0.0
		for _, n := range numbers {
			sum += n // Просто прибавляем каждое число к переменной sum.
		}
		return sum, nil

	case "AVG":
		sum := 0.0
		for _, n := range numbers {
			sum += n
		} // Делим сумму на количество элементов в списке.
		return sum / float64(len(numbers)), nil

	case "MED":
		sort.Float64s(numbers) // Сначала СТРОГО сортируем числа по возрастанию.
		n := len(numbers)
		if n%2 == 1 {
			// Если количество чисел нечетное, берем центральное.
			return numbers[n/2], nil
		} // Если четное — берем среднее между двумя центральными числами.
		return (numbers[n/2-1] + numbers[n/2]) / 2, nil

	default:
		return 0, fmt.Errorf("неизвестная операция: %s", operation)
	}
}

func main() {
	var op, input string // Объявляем переменные для операции и строки ввода.

	fmt.Print("Введите операцию (AVG, SUM, MED): ")
	fmt.Scanln(&op) // Считываем ввод пользователя в переменную op.

	fmt.Print("Введите числа через запятую: ")
	var line string
	fmt.Scanln(&line) // Считываем строку с числами.
	input = line
	// Вызываем нашу функцию и получаем результат и ошибку.
	result, err := calculate(op, input)
	if err != nil { // Если функция вернула ошибку, печатаем её и прекращаем работу (return).
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	// Если всё хорошо, печатаем результат с двумя знаками после запятой (%.2f).
	fmt.Printf("Результат: %.2f\n", result)
}
