package main

import (
	"fmt"
	"log"
	"net/http"
	"simple/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("server is started...")

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("listening at port 4000....")

}
