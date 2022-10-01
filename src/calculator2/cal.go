package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("숫자를 입력하세요.")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n1, _ := strconv.Atoi(line)
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n2, _ := strconv.Atoi(line)

	fmt.Printf("입력하신 숫자는 %d, %d 입니다.\n", n1, n2)

	fmt.Println("연산자를 입력하세요")
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	switchCaseUse(n1, n2, line)
	// ifUse(n1, n2, line)
}

func ifUse(n1 int, n2 int, op string) {
	if op == "+" {
		fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	} else if op == "-" {
		fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	} else if op == "*" {
		fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
	} else if op == "/" {
		fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	} else {
		fmt.Println("잘못 입력하셨습니다.")
	}
}

func switchCaseUse(n1 int, n2 int, op string) {
	switch op {
	case "+":
		fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	case "-":
		fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	case "*":
		fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
	case "/":
		fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	default:
		fmt.Println("잘못 입력하셨습니다.")
	}
}
