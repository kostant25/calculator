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
	default:
		panic(errors.New("Invalid operation"))
	}

	if flag {
		fmt.Println("Rome: " + strconv.Itoa(result))
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
	for i := len(str) - 1; i >= 0; i-- {
		switch str[i] {
		case 73: // I
			num += 1
		case 86: // V
			if i > 0 {
				if str[i-1] == 73 {
					num += 4
					i--
				}
			} else {
				num += 5
			}
		case 88: // X
			if i > 0 {
				if str[i-1] == 73 {
					num += 9
					i--
				}
			} else {
				num += 10
			}
		}
	}
	return
}

func arabicToRome(num int) (str string) {

	return
}
