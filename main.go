package main

import "fmt"

func main() {
	s := print("Артём", 44)
	fmt.Println(s)
}
func print(name string, age int) string {

	res := fmt.Sprintf("Меня зовут %s и мне %d года", name, age)

	return res
}

// Задача
// Создать функцию print, которая принимает в себя имя пользователя и его возраст
// Функция print возвращает строку формата: "Меня зовут name и мне age лет"
