package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Как тебя зовут?: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	fmt.Print("Сколько тебе лет: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	fmt.Println("меня зовут", input, "и мне", input1, "года")
}
