package main

import (
	"fmt"
	//"strconv"
)

func main() {
	//s :=  [] int{1,2,3,4}
	//add(1, 2, 3, 4, 5)
	in()

}

func add(s ...int) {
	var a = 1
	for i := 1; i < len(s); i++ {

		a = a * i

	}
	fmt.Println(a)
}

func in() {
	co := []int{1, 2, 3, 4, 5}
	var index int = 2

	//re := append(co[1:3])
	re := append(co[:index], co[index+1:]...)
	fmt.Println(re)
}
