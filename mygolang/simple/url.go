package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("welcome to the Url Program")
	newurl := &url.URL{
		Scheme: "https",
		Path:   "/tutcss",
		Host:   "lco.dev",
	}
	fmt.Println(newurl.String())

}
