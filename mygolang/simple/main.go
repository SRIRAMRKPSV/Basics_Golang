package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("welcome to user input")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("the user input is \n")

	input, _ := reader.ReadString('\n')
	fmt.Println("the value is,", input)
}
