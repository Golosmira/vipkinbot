package main

import "fmt"

func main() {
	res := sum(5, 2) + sum(7, 9) + sum(9, 13) + sum(1, 9)

	fmt.Println(res)
}

func sum(a int, b int) int {
	s := (a + b) * 2

	return s
}

// Задача
// Создать функцию print, которая принимает в себя имя пользователя и его возраст
// Функция print возвращает строку формата: "Меня зовут name и мне age лет"
