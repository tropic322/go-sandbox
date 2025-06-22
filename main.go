package main

import (
	"bufio"
	"fmt"
	"os"
	bracketService "task-1/services/bracket"
	stringService "task-1/services/string"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Ведите номер задания: ")
	reader.Scan()
	input := reader.Text()

	switch input {
	case "1":
		stringService.GetInputStringLength()
	case "2":
		stringService.GetVowelLetters()
	case "3":
		stringService.Capitalize()
	case "4":
		bracketService.CheckBrackets()
	default:
		fmt.Println("Введенное число не соответствует номеру ни одного задания")
	}
}
