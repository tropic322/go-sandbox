package bracket

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func CheckBrackets() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Ведите строку: ")
	reader.Scan()
	input := reader.Text()
	pattern := `\(`
	reLeft := regexp.MustCompile(pattern)
	leftBracketsAmount := len(reLeft.FindAllString(input, -1))

	pattern = `\)`
	reRight := regexp.MustCompile(pattern)
	rightBracketsAmount := len(reRight.FindAllString(input, -1))

	if leftBracketsAmount != rightBracketsAmount {
		fmt.Printf("вывод: Скобки расставлены неправильно, %v открывающиеся, %v закрывающиеся", leftBracketsAmount, rightBracketsAmount)
	}
}
