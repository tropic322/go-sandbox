package getters

import (
	"strconv"
	"strings"
)

func GetNumbers(inputString string) []int {
	// Убираем все разделители кроме цифр
	cleaned := strings.Map(func(r rune) rune {
		if r == ',' {
			return ' '
		}
		if r >= '0' && r <= '9' || r == ' ' {
			return r
		}
		return -1
	}, inputString)

	parts := strings.Fields(cleaned)

	var numbers []int
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	return numbers
}

func SortNumbers(numbers []int) {
	n := len(numbers)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j+1] > numbers[j] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}

func GetAverage(numbers []int) int {
	n := len(numbers)
	var sum int
	for i := 0; i < n; i++ {
		sum += numbers[i]
	}
	return sum / n
}
