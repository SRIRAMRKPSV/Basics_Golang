package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("welcome to the mod program")
	mod()
	mods()
	r := mux.NewRouter()
	r.HandleFunc("/", final).Methods("GET")

	http.ListenAndServe(":3000", r)
	log.Fatal(http.ListenAndServe(":3000", r))

}
func mods() {

	fmt.Println("hi sriram")

}
func mod() {
	fmt.Println("vidhya lakshmi is the mental")
}

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1><marquee>VidhyaLakshmi is the bad girl<marquee></h1>"))

}
