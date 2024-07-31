package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Возникла ошибка:", r)
		}
	}()

	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("Введите математическое выражение (например, 1+1 или I+I):")
		scanner.Scan()
		userInput := scanner.Text()

		result := evaluateExpression(userInput)
		fmt.Println("Результат:", result)
	}
}

func evaluateExpression(input string) string {
	operators := []string{"+", "-", "*", "/"}

	for _, operator := range operators {
		parts := strings.Split(input, operator)
		if len(parts) == 2 {
			leftStr := strings.TrimSpace(parts[0])
			rightStr := strings.TrimSpace(parts[1])

			leftIsRoman := isRoman(leftStr)
			rightIsRoman := isRoman(rightStr)

			if leftIsRoman != rightIsRoman {
				panic("нельзя смешивать римские и арабские цифры в одном выражении")
			}

			var left, right int
			var err error

			if leftIsRoman {
				left, err = romanToInt(leftStr)
				if err != nil {
					panic(err)
				}
				right, err = romanToInt(rightStr)
				if err != nil {
					panic(err)
				}
			} else {
				left, err = strconv.Atoi(leftStr)
				if err != nil {
					panic(err)
				}
				right, err = strconv.Atoi(rightStr)
				if err != nil {
					panic(err)
				}
			}

			if left > 10 || right > 10 {
				panic("числа не должны быть больше 10")
			}

			var result int
			switch operator {
			case "+":
				result = left + right
			case "-":
				result = left - right
			case "*":
				result = left * right
			case "/":
				if right == 0 {
					panic("деление на ноль")
				}
				result = left / right
			}

			if leftIsRoman {
				if result <= 0 {
					panic("результат римского числа не может быть отрицательным или нулевым")
				}
				return intToRoman(result)
			}
			return strconv.Itoa(result)
		}
	}

	panic("не удалось распознать выражение")
}

func isRoman(s string) bool {
	for _, c := range s {
		if c != 'I' && c != 'V' && c != 'X' {
			return false
		}
	}
	return true
}

func romanToInt(s string) (int, error) {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	n := len(s)
	total := 0

	for i := 0; i < n; i++ {
		value, ok := romanMap[rune(s[i])]
		if !ok {
			return 0, fmt.Errorf("неверное римское число: %s", s)
		}

		if i < n-1 && value < romanMap[rune(s[i+1])] {
			total -= value
		} else {
			total += value
		}
	}

	return total, nil
}

func intToRoman(num int) string {
	val := []int{50, 10, 9, 5, 4, 1}
	sym := []string{"L", "X", "IX", "V", "IV", "I"}

	var result strings.Builder
	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			result.WriteString(sym[i])
		}
	}
	return result.String()
}
