package main

import "fmt"

func main() {
	defer fmt.Println("this will execute in  the last line ")
	fmt.Println("the program is based on defer")
	defers()

}

func defers() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
