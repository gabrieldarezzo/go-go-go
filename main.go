package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	number1, _ := RetrieveNumberByIo()
	number2, _ := RetrieveNumberByIo()

	sumValue, _ := Sum(number1, number2)

	//subValue, err := Subtraction(5, 3)

	//multiValue, err := Multi(5, 3)

	//divideValue, err := Divide(2, 0)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	fmt.Println(sumValue)
	//fmt.Println(divideValue)

}

func Sum(num1, num2 float64) (float64, error) {
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

func RetrieveNumberByIo() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number: ")
	text, _ := reader.ReadString('\n')
	result := strings.Split(strings.TrimSpace(text), "x")

	return strconv.ParseFloat(result[0], 64)
}
