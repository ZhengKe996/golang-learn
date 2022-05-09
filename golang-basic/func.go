package main

import "fmt"

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return -1, fmt.Errorf("unsupported operation:" + op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func main() {
	fmt.Println(eval(3, 4, "x"))
	fmt.Println(eval(3, 4, "+"))
}
