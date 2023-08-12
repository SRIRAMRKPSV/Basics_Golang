package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	fmt.Println("welcome to web request Port")
	//requests()
	//responses()
	forms()

}
func requests() {
	const ports = "http://localhost:4000/get"
	response, err := http.Get(ports)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}

func responses() {
	const posts = "http://localhost:3000/post"
	contents := strings.NewReader(`
	{
       "name" : "sriram"
	   
	}
	
	`)
	responsess, err := http.Post(posts, "application/json", contents)
	if err != nil {
		panic(err)
	}
	defer responsess.Body.Close()
	content, err := ioutil.ReadAll(responsess.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
func forms() {
	const urls = "http://localhost:3000/postform"
	datas := &url.Values{}
	datas.Add("name", "sriram")
	datas.Add("email", "sriram@google.in")

	respon, err := http.PostForm(urls, *datas)
	if err != nil {
		panic(err)
	}
	defer respon.Body.Close()
	contents, err := ioutil.ReadAll(respon.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(contents))
}
