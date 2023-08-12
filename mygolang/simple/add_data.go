package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("welcome to the web process using json")
	sites()

}
func sites() {
	data := []course{
		{"fullstack", "ineuron.ai", "abc@123", []string{"chennai", "anna_nagar"}},
		{"pyhon", "lco.ai", "abc@123", nil},
		{"machine_learning", "course.ai", "abc@123", []string{"bangalore", "mysore"}},
	}
	conv, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		panic(err)

	}

	fmt.Printf(" %s", conv)
}

type course struct {
	Name     string   `json:"course_name"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Address  []string `json:",omitempty"`
}
