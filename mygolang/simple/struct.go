package main

import "fmt"

func main() {

	fmt.Println("welcome to struct Program")
	sr := User{"sriram", 9080504893, "sriram@gmail.com", true, 23, 173.4}
	fmt.Printf("The values in the struct are :%+v", sr)

}

type User struct {
	Name         string
	Phone_number int
	Email        string
	Status       bool
	Age          int
	Height       float32
}
