// After line 23 there is  a goto statement
package main

import "fmt"

func main() {
	//var a int
	//a = 3
	//n := []int{1, 2, 3, 4}
	//fmt.Println("enter the number")
	//fmt.Scanln(&n)  // user input
	//fmt.Println(" the number is :", n)
	//var result string
	//for i := 1; i <= len(n); i++ {
	//if a == n[i] {
	//result = "the value is found"
	//} else {
	//	result = "the value is not fond"
	//}
	//fmt.Println(result)

	//}

	na := 2
	for na < 10 {
		if na == 1 {
			fmt.Println("the value is one")
			na++
		}
		if na == 3 {
			fmt.Println("the value is three")
			na++
		}
		if na == 2 {
			fmt.Println("the value is two")
			na++
		}
		if na == 4 {
			fmt.Println("the value is four")
			goto sri
		}

	}
sri:
	fmt.Println("this is my function")
}
