package main

import (
	"fmt"
)

func main() {
	fmt.Println("this is the function Program")
	new := User{"sriram", 23, 9080504893}
	fmt.Printf("user details%+v \n", new)
	new.details()

}

type User struct {
	Name  string
	Age   int
	phone int
}

func (u User) details() {
	u.phone = 9080504983
	fmt.Println("new Phone nuber:", u.phone)

}
