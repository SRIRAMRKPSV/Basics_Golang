package main

import "fmt"

func main() {
	fmt.Println("welcome to the pointer program")

	var num = 15

	var ptr = &num
	fmt.Println("the value of the pointer", ptr)
	fmt.Println("the value of the pointer", *ptr)

	*ptr = *ptr * 2

	fmt.Println("the value of the pointer", num)

}
