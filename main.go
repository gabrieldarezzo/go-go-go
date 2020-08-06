package main

import (
	"fmt"
)

func main() {
	//sumValue, err := Sum(1, 2)

	subValue, err := Subtraction(5, 3)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(subValue)
}

func Sum(num1, num2 int) (int, error) {
	return num1 + num2, nil
}

func Subtraction(num1, num2 float64) (float64, error) {
	return num1 - num2, nil
}
