package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Input expression:")
	operand1, operand2, operation := read()

	switch operation {
	case "+":
		fmt.Println(operand1 + operand2)
	case "-":
		fmt.Println(operand1 - operand2)
	case "*":
		fmt.Println(operand1 * operand2)
	case "/":
		fmt.Println(operand1 / operand2)
	}

}

func read() (o1 int, o2 int, op string) {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	expression := strings.Split(strings.Trim(line, "\n"), " ")
	o1 = parseToInt(expression[0])
	o2 = parseToInt(expression[2])
	op = expression[1]
	return
}

func parseToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return num
}
