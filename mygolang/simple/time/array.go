package main

import "fmt"

func main() {

	a := []string{"s", "r", "i", "r", "a", "m"}
	var ind int = 2
	a = append(a[:ind], a[ind+1:]...)
	fmt.Println(a)
}
