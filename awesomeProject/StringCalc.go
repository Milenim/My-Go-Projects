package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var userInput []string
	seps := []string{"+", "/", "-", "*"}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	var sep string
	for _, sep = range seps {
		if strings.Contains(text, sep) {
			userInput = strings.Split(text, sep)
			break
		}
	}

	str1 := strings.Trim(userInput[0], "\" ")
	secondOperand := strings.TrimSpace(userInput[1])

	switch sep {
	case "+":
		str2 := strings.Trim(secondOperand, "\"")
		result := str1 + str2
		fmt.Printf("\"%s\"\n", result)
	case "*":
		parseInt, err := strconv.Atoi(secondOperand)
		if err != nil || parseInt < 1 || parseInt > 10 {
			panic("Неверный множитель: ожидается целое число от 1 до 10")
		}
		result := strings.Repeat(str1, parseInt)
		fmt.Printf("\"%s\"\n", result)
	case "/":
		parseInt, err := strconv.Atoi(secondOperand)
		if err != nil || parseInt < 1 || parseInt > 10 {
			panic("Неверный делитель: ожидается целое число от 1 до 10.")
		}
		if len(str1) < parseInt {
			panic("Длина строки короче делителя")
		}
		result := str1[:len(str1)/parseInt]
		fmt.Printf("\"%s\"\n", result)
	case "-":
		str2 := strings.Trim(secondOperand, "\"")
		if !strings.Contains(str1, str2) {
			panic("Вычитаемое не найдено в строке")
		}
		result := strings.Replace(str1, str2, "", 1)
		fmt.Printf("\"%s\"\n", result)
	default:
		panic("Неподдерживаемая операция")
	}
}
