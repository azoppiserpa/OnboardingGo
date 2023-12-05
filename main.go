package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the first integer: ")
	firstIntInput, _ := reader.ReadString('\n')
	firstInt, _ := strconv.Atoi(strings.TrimSpace(firstIntInput))

	fmt.Print("Enter the second integer: ")
	secondIntInput, _ := reader.ReadString('\n')
	secondInt, _ := strconv.Atoi(strings.TrimSpace(secondIntInput))

	fmt.Print("Enter the operation (a or s): ")
	operationInput, _ := reader.ReadString('\n')
	operation := strings.TrimSpace(operationInput)

	var result int

	switch operation {
	case "a":
		result = add(firstInt, secondInt)
	case "s":
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
