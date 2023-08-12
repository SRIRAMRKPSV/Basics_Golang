package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("this is the program to create and write in a file")
	content := "this is the program to create and write in a file"
	file, err := os.Create("./myfirst.txt")

	if err != nil {
		panic(err)
	}
	insert, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println(insert)
	defer file.Close()
	readfile("./myfirst.txt")

}
func readfile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
