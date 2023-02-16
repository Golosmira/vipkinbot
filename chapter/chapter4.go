package main

import (
	"fmt"
	"strings"
)

// Operation function // Функция управления
type Operate func(int, int) int

// Structure of operands and the operator // Структура операндов и оператор
type Expression struct {
	X, Y      int
	Operation Operate
}

// Map of single digits // Карта однозначных чисел
var singledigits = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
}

// Map of operators "+" "-" and funcs Карта операторов "+" "-" и функций
var operators = map[string]Operate{
	"+": func(x, y int) int { return x + y },
	"-": func(x, y int) int { return x - y },
}

// Filling Expression structure Заполняющая структура выражения
func (exp Expression) FillingExpression(stringarr []string) Expression {
	for _, elem := range stringarr {
		_, ok := singledigits[elem]

		if ok {
			exp.X = singledigits[stringarr[0]]
			exp.Y = singledigits[stringarr[2]]
		} else {
			exp.Operation = operators[elem]
		}

	}
	return Expression{exp.X, exp.Y, exp.Operation}
}

// Preparing input condition with trim spaces Подготовка входного условия с помощью пробелов для обрезки
func PreparingInputCondition(condition string) []string {
	stringArr := []string{}
	conditionArr := strings.Split(condition, "")

	for _, str := range conditionArr {
		if str != " " {
			stringArr = append(stringArr, str)
		}
	}
	return stringArr
}

func main() {
	inputCondition := "1+ 1" // Второе условие "2-1 + 2"
	preparedCondition := PreparingInputCondition(inputCondition)
	fmt.Println("Prepared condition: ", PreparingInputCondition(inputCondition))

	expression := Expression{}

	completeExpression := expression.FillingExpression(preparedCondition)

	result := Expression{
		X:         completeExpression.X,
		Y:         completeExpression.Y,
		Operation: completeExpression.Operation,
	}

	fmt.Println("Result of operation: ", result.Operation(result.X, result.Y))

}
