package main

import (
	"fmt"
)

func main() {
	n := 5
	a := 1
	for i := 1; i < n; i++ {
		a = a * i

	}
	fmt.Println(" ", a)
}
