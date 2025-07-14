package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	slice := getUserInput()

	switch getOperationInput() {
	case "AVG":
		fmt.Println("AVG = ", calculateAVG(slice))
	case "SUM":
		fmt.Println("SUM = ", calculateSUM(slice))
	case "MED":
		fmt.Println("MED = ", calculateMED(slice))
	default:
	}
}

func getUserInput() []float64 {
	var numbers []float64

	fmt.Println("Введите числа через запятую: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	input = strings.TrimSpace(input)
	parts := strings.Split(input, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		num, err := strconv.ParseFloat(part, 64)
		if err != nil {
			fmt.Println("Ошибка в одном из числе")
			continue
		}
		numbers = append(numbers, num)

	}
	return numbers
}

func getOperationInput() string {
	var operation string
	fmt.Println("Введите название операции (AVG, SUM, MED)")
	for {
		fmt.Scan(&operation)
		if operation == "AVG" || operation == "MED" || operation == "SUM" {
			return operation
		}
		fmt.Println("Попробуйте AVG, SUM, MED")
	}
}

func calculateAVG(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0.0
	}

	return calculateSUM(numbers) / float64(len(numbers))
}

func calculateSUM(numbers []float64) float64 {
	var sum float64
	if len(numbers) == 0 {
		return 0.0
	}

	for _, num := range numbers {
		sum += num
	}
	return sum
}

func calculateMED(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0.0
	}

	sort.Slice(numbers, func(i, j int) bool { return numbers[i] < numbers[j] })

	if len(numbers)%2 == 0 {
		return (numbers[len(numbers)/2-1] + numbers[len(numbers)/2]) / 2
	}
	return numbers[len(numbers)/2]
}
