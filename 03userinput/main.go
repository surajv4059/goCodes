package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the reading of out pizza: ")

	//comma ok or err ok
	input, _ := reader.ReadString(('\n'))
	fmt.Println("Thanks for rating, ", input)

}
