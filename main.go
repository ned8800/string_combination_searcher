package main

import (
	"fmt"
	"strconv"
)

const (
	givenString = "9876543210"
	targetSum   = 200
)

// Рекурсивная функция для поиска комбинаций, которые в сумме дают результат 200
// index отвечает за текущее положение каретки в строке
// currentExpression отвечает за оставшуюся часть строки, которую нужно проверить
// currentSum отвечает за текущую сумму от всех предыдущих шагов
func stringCombinationSearcher(index int, currentExpression string, currentSum int) {
	if index > len(givenString) || index < 0 || len(givenString) < 1 {
		fmt.Println("Индекс не должен быть больше длины строки и не должен быть меньше нуля")
		return
	}

	if index == len(givenString) {
		if currentSum != targetSum {
			return
		}

		fmt.Printf("%s=%d\n", currentExpression, currentSum)
	}

	for i := index + 1; i <= len(givenString); i++ {
		currentNumStr := givenString[index:i]
		currentNum, err := strconv.Atoi(currentNumStr)
		if err != nil {
			fmt.Printf("Строка не состоит только из цифр. error: %v", err)
			return
		}

		if index == 0 {
			stringCombinationSearcher(i, currentNumStr, currentNum)
		} else {
			stringCombinationSearcher(i, currentExpression+"+"+currentNumStr, currentSum+currentNum)
			stringCombinationSearcher(i, currentExpression+"-"+currentNumStr, currentSum-currentNum)
		}

	}

}

func main() {
	fmt.Printf("Комбинации для строки \"%s\", которые в результате дают %d:\n", givenString, targetSum)
	stringCombinationSearcher(0, "", 0)
}
