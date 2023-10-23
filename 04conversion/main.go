package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("welcome to our pizza app")
	fmt.Println("pelase rate our pizza between 1 and 5")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString(('\n'))
	fmt.Println("this is input ", input)
	numRating, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println(err)
	}
}
