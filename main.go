package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	result := add(5, 3)
	fmt.Println("Addition result:", result)

	result = subtract(10, 7)
	fmt.Println("Subtraction result:", result)
}
package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the first integer: ")
	firstIntInput, _ := reader.ReadString('\n')
	firstInt, _ := strconv.Atoi(firstIntInput)

	fmt.Print("Enter the second integer: ")
	secondIntInput, _ := reader.ReadString('\n')
	secondInt, _ := strconv.Atoi(secondIntInput)

	fmt.Print("Enter the operation (+ or -): ")
	operationInput, _ := reader.ReadString('\n')
	operation := operationInput[:len(operationInput)-1]

	var result int
	switch operation {
	case "+":
		result = add(firstInt, secondInt)
	case "-":
		result = subtract(firstInt, secondInt)
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Println("Result:", result)
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}
