package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Input expression:")
	operand1, operand2, operation, flag := read()
	var result int

	switch operation {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		result = operand1 / operand2
	}

	if flag {
		fmt.Println(arabicToRome(result))
	} else {
		fmt.Println(result)
	}

}

func read() (o1 int, o2 int, op string, flag bool) {
	var flag1 bool
	var flag2 bool
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	expression := strings.Split(strings.Trim(line, "\n"), " ")
	o1, flag1 = parseToInt(expression[0])
	o2, flag2 = parseToInt(expression[2])
	op = expression[1]

	if flag1 == flag2 {
		flag = flag1
	} else {
		err := errors.New("Different numeric alphavit")
		log.Fatal(err)
	}

	return
}

func parseToInt(s string) (int, bool) {
	var flag bool
	num, err := strconv.Atoi(s)
	if err != nil {
		flag = true
		num = romeToArabic(s)
	}
	return num, flag
}

func romeToArabic(str string) (num int) {

	return
}

func arabicToRome(num int) (str string) {

	return
}
