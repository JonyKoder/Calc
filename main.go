package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var romeMap = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func arabToRome(number int) string {
	if number <= 0 {
		return "Результат не может быть меньше одного"
	}
	romeNumbers := []struct {
		Value  int
		Symbol string
	}{
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result strings.Builder

	for _, rn := range romeNumbers {
		for number >= rn.Value {
			result.WriteString(rn.Symbol)
			number -= rn.Value
		}
	}

	return result.String()
}

func romeToArab(rome string) (int, error) {
	result := 0
	var prevValue int

	for i := len(rome) - 1; i >= 0; i-- {
		value := romeMap[string(rome[i])]
		if value < prevValue {
			result -= value
		} else {
			result += value
		}
		prevValue = value
	}

	return result, nil
}

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			fmt.Println("Деление на ноль")
			return 0
		}
		return a / b
	default:
		fmt.Println("Неподдерживаемая операция")
		return 0
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите выражение в формате 'число операция число', например: 2 + 3 (для выхода нажмите Q)")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToUpper(input) == "Q" {
			break
		}

		parts := strings.Fields(input)

		if len(parts) != 3 {
			panic("Некорректный ввод. Пожалуйста, введите выражение в формате 'число операция число'")
		}

		num1Type := getNumType(parts[0])
		num2Type := getNumType(parts[2])

		a, err := convertToNumber(parts[0], num1Type)
		if err != nil {
			fmt.Println("Ошибка ввода первого числа:", err)
			continue
		}

		b, err := convertToNumber(parts[2], num2Type)
		if err != nil {
			fmt.Println("Ошибка ввода второго числа:", err)
			continue
		}
		if a > 10 || b > 10 {
			panic("Число не должно быть больше 10")
		}
		result := calculate(a, b, parts[1])
		if result < 0 {
			panic("Результат не может быть меньше нуля")
		}
		if (num1Type == "roman" && num2Type == "arabic") || (num2Type == "roman" && num1Type == "arabic") {
			panic("Используются одновременно разные системы счисления")
		}
		if num1Type == "roman" || num2Type == "roman" {
			resultRoman := arabToRome(result)
			fmt.Println("Результат в римской системе:", resultRoman)
		} else {
			fmt.Println("Результат в арабской системе:", result)
		}
	}
}

func getNumType(num string) string {
	match, _ := regexp.MatchString("^[IVXLCDMivxlcdm]+$", num)
	if match {
		arabNum, _ := romeToArab(num)

		if arabNum > 10 {
			panic("Число не должно быть больше 10")
		}
		return "roman"
	}
	return "arabic"
}
func convertToNumber(num string, numType string) (int, error) {
	if numType == "roman" {
		return romeToArab(num)
	}
	return strconv.Atoi(num)
}
