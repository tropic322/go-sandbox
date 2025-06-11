package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Ведите числа: ")
	reader.Scan()
	input := reader.Text()

	// Убираем все разделители кроме цифр
	cleaned := strings.Map(func(r rune) rune {
		if r == ',' {
			return ' '
		}
		if r >= '0' && r <= '9' || r == ' ' {
			return r
		}
		return -1
	}, input)

	parts := strings.Fields(cleaned)

	var numbers []int
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err == nil {
			numbers = append(numbers, num)
		}
	}

	fmt.Println(numbers)
}
