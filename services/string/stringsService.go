package string

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func GetInputStringLength() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Ведите строку: ")
	reader.Scan()
	input := reader.Text()
	fmt.Printf("Количество символов в строке = %v", len(input))
}

func GetVowelLetters() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Ведите строку: ")
	reader.Scan()
	input := reader.Text()
	pattern := `[а-яёА-ЯЁ]`
	re := regexp.MustCompile(pattern)
	fmt.Printf("Количество символов в строке = %v", len(re.FindAllString(input, -1)))
}

func Capitalize() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Ведите строку: ")
	reader.Scan()
	input := reader.Text()
	words := strings.Fields(input)
	for i, word := range words {
		runes := []rune(word)
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
			for j := 1; j < len(runes); j++ {
				runes[j] = unicode.ToLower(runes[j])
			}
			words[i] = string(runes)
		}
	}

	fmt.Printf("Итоговая строка - %v", strings.Join(words, " "))
}
