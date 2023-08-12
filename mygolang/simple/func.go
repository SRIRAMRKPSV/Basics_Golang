package main

import "fmt"

func main() {
	fmt.Println("welcome to func program")

	add := simple(1, 2, 3, 4)
	fmt.Println(add)

}
func simple(values ...int) int {
	var a int
	for i := 0; i < len(values); i++ {

		a = a + values[i]

	}
	return a
}
