package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("WElcome to my Program")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the user input")

	input, err := reader.ReadString('\n')

	fmt.Println("thanks for the input:", input)
	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The number is added", num+1)

	}

}
