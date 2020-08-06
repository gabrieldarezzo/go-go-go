package main

import (
	"errors"
	"fmt"
	"bufio"
	"os"
)



func main() {

	

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ") 
	text, _ := reader.ReadString('\n')

	fmt.Println("Vc digitou:") 
	fmt.Println(text)

	sumValue, err := Sum(1, 2)

	//subValue, err := Subtraction(5, 3)

	//multiValue, err := Multi(5, 3)

	//divideValue, err := Divide(2, 0)

	
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sumValue)
	//fmt.Println(divideValue)
	
}

func Sum(num1, num2 int) (int, error) {
	return num1 + num2, nil
}

func Subtraction(num1, num2 float64) (float64, error) {
	return num1 - num2, nil
}

func Multi(num1, num2 float64) (float64, error) {
	return num1 * num2, nil
}

func Divide(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("Wrong value")
	}
	return num1 / num2, nil
}
