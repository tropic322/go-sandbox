package main

import (
	"bufio"
	"fmt"
	"os"
	NG "task-1/services/getters"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите числа: ")
	reader.Scan()
	input := reader.Text()

	numbers := NG.GetNumbers(input)
	fmt.Println("Массив, который был вычитан: ", numbers)
	NG.SortNumbers(numbers)
	fmt.Println("Отсортированные элементы: ", numbers)
	fmt.Println("Самое большое число: ", numbers[0])
	fmt.Println("Самое маленькое число: ", numbers[len(numbers)-1])
}
