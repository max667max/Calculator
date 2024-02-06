package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Введите выражение: ")
	input, _ := reader.ReadString('\n')
	result, roman, err := calculate(input)
	if err != nil {
		log.Fatal(err)
	}
	if !roman {
		fmt.Println("Результат операции с арабскими числами: ", result)
	} else {
		fmt.Println("Резулльтат операции с римскими чисами: ", intToRoman(result))
	}

}

func calculate(input string) (int, bool, error) {
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		panic("Неверный формат ввода") // разделение входящей строки на операнды и оператор
	}

	operand1, err := parseOperand(parts[0]) //парсинг операндов
	if err != nil {
		return 0, true, err
	}

	operand2, err := parseOperand(parts[2])
	if err != nil {
		return 0, true, err
	}

	operator := parts[1]

	// Проверка на одновременное использование разных систем счисления
	if (isRoman(parts[0]) && isArabic(parts[2])) || (isArabic(parts[0]) && isRoman(parts[2])) {
		panic("Вы одновременно использовали разные системы счисления")
	}

	// Выполнение операции
	result := 0
	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		result = operand1 / operand2
	default:
		panic("Недопустимый оператор")
	}

	// Проверка отрицательного результата для римских чисел
	if isRoman(parts[0]) && isRoman(parts[2]) && result <= 0 {
		panic("Операция с римскими числами дала нулевой или отрицательный результат")
	}

	return result, isRoman(parts[0]), nil
}

// Парсин операндов
func parseOperand(operand string) (int, error) {
	if isRoman(operand) {
		return parseRoman(operand)
	}
	return parseArabic(operand)
}

// проверка, является ли операнд римским числом
func isRoman(operand string) bool {
	if strings.Contains(operand, "I") || strings.Contains(operand, "V") || strings.Contains(operand, "X") {
		return true
	}
	return false
}

// парсинг римских чисел
func parseRoman(operand string) (int, error) {

	romanNums := map[rune]int{'I': 1, 'V': 5, 'X': 10}

	result := 0

	prevVal := 0

	for i := 0; len(operand)-1 >= i; i++ {
		vale := romanNums[rune(operand[i])]
		if vale > prevVal {
			result = vale - prevVal
		} else {
			result += vale
		}
		prevVal = vale
	}

	return result, nil
}

// проверка, является ли операнд арабским числом
func isArabic(operand string) bool {
	if strings.Contains("123456789", string(operand[0])) {
		return true
	}
	return false
}

// парсинг арабских чисел
func parseArabic(operand string) (int, error) {

	if len(operand) == 1 || (len(operand) == 3 && !strings.Contains("1234567890", string(operand[1]))) {

		arabicNums := map[rune]int{
			'1': 1,
			'2': 2,
			'3': 3,
			'4': 4,
			'5': 5,
			'6': 6,
			'7': 7,
			'8': 8,
			'9': 9,
		}

		return arabicNums[rune(operand[0])], nil

	}

	if (len(operand) == 2 || (len(operand) == 4 && !strings.Contains("1234567890", string(operand[2])))) && operand[0] == '1' && operand[1] == '0' {
		return 10, nil
	} else {
		panic("Один из оперрандоов выходит из заданного диапазона")
	}

}

func intToRoman(num int) string {

	romanNumerals := map[int]string{
		1:   "I",
		4:   "IV",
		5:   "V",
		9:   "IX",
		10:  "X",
		40:  "XL",
		50:  "L",
		90:  "XC",
		100: "C",
	}

	result := ""
	values := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}

	for _, value := range values {
		for num >= value {
			result += romanNumerals[value]
			num -= value
		}
	}

	return result

}
