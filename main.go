package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите операцию (AVG, SUM, MED): ")
	operation, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	operation = strings.TrimSpace(operation)

	fmt.Print("Введите числа через запятую: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)

	numbers, err := parseNumbers(input)
	if err != nil {
		log.Fatal(err)
	}

	result, err := calculate(operation, numbers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Результат: %v\n", result)
}

// parseNumbers разбирает строку вида "1, 2, 3.5" в срез float64
func parseNumbers(input string) ([]float64, error) {
	if input == "" {
		return nil, fmt.Errorf("ввод пуст")
	}

	parts := strings.Split(input, ",")
	var numbers []float64

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}
		num, err := strconv.ParseFloat(trimmed, 64)
		if err != nil {
			return nil, fmt.Errorf("некорректное число: %s", trimmed)
		}
		numbers = append(numbers, num)
	}

	if len(numbers) == 0 {
		return nil, fmt.Errorf("не найдено ни одного корректного числа")
	}

	return numbers, nil
}

// calculate выполняет операцию над числами и возвращает результат или ошибку
func calculate(op string, nums []float64) (float64, error) {
	switch op {
	case "SUM":
		return sum(nums), nil
	case "AVG":
		return avg(nums), nil
	case "MED":
		return med(nums), nil
	default:
		return 0, fmt.Errorf("неизвестная операция: %s. Используйте AVG, SUM или MED", op)
	}
}

func sum(nums []float64) float64 {
	var total float64
	for _, value := range nums {
		total += value
	}
	return total
}

func avg(nums []float64) float64 {
	return sum(nums) / float64(len(nums))
}

func med(nums []float64) float64 {
	sorted := make([]float64, len(nums))
	copy(sorted, nums)
	sort.Float64s(sorted)
	n := len(sorted)
	if n%2 == 0 {
		return (sorted[n/2-1] + sorted[n/2]) / 2
	}
	return sorted[n/2]
}
