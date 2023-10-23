package main

import "fmt"

func main() {
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallint uint8 = 255
	fmt.Println(smallint)
	fmt.Printf("variable is of type: %T \n", smallint)

	var smallfloat float64 = 255.865435666666666624534534536
	fmt.Println(smallfloat)
	fmt.Printf("variable is of type: %T \n", smallfloat)

}
