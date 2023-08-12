package main

import "fmt"

func main() {
	var a int
	var result string
	fmt.Print("Enter the number\n")
	fmt.Scanln(&a)
	if a%2 == 0 {
		result = "Even"
	} else {
		result = "Odd"
	}
	fmt.Println(result)
}
