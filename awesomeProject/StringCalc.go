package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	sep, startIndex := findOperatorOutsideQuotes(text)
	if sep == "" {
		panic("Неверный оператор операции")
	}

	userInput := []string{
		strings.TrimSpace(text[:startIndex]),
		strings.TrimSpace(text[startIndex+1:]),
	}

	str1 := strings.TrimSpace(strings.Trim(userInput[0], "\""))
	secondOperand := strings.TrimSpace(userInput[1])

	if len(str1) > 10 {
		panic("Длина строки больше 10 символов")
	}

	var result string

	switch sep {
	case "+":
		str2 := strings.Trim(secondOperand, "\"")
		result = str1 + str2
	case "*":
		parseInt, err := strconv.Atoi(secondOperand)
		if err != nil || parseInt < 1 || parseInt > 10 {
			panic("Неверный множитель: ожидается целое число от 1 до 10")
		}
		result = strings.Repeat(str1, parseInt)
	case "/":
		parseInt, err := strconv.Atoi(secondOperand)
		if err != nil || parseInt < 1 || parseInt > 10 {
			panic("Неверный делитель: ожидается целое число от 1 до 10")
		}
		if len(str1) < parseInt {
			panic("Длина строки короче делителя")
		}
		result = str1[:len(str1)/parseInt]
	case "-":
		str2 := strings.Trim(secondOperand, "\"")
		if !strings.Contains(str1, str2) {
			panic("Вычитаемое не найдено в строке")
		}
		result = strings.Replace(str1, str2, "", 1)
	default:
		panic("Неподдерживаемая операция")
	}

	if len(result) > 40 {
		result = result[:40] + "..."
	}

	fmt.Printf("\"%s\"\n", result)
}

func findOperatorOutsideQuotes(text string) (string, int) {
	seps := []string{"+", "/", "-", "*"}
	inQuotes := false

	for i, char := range text {
		switch char {
		case '"':
			inQuotes = !inQuotes
		default:
			if !inQuotes {
				for _, sep := range seps {
					if string(char) == sep {
						return sep, i
					}
				}
			}
		}
	}
	return "", -1
}
