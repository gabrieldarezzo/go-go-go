package main

import (
	"fmt"
)

func main() {
	sumValue, err := Sum(1, 2)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sumValue)
}

func Sum(num1, num2 int) (int, error) {
	return num1 + num2, nil
}