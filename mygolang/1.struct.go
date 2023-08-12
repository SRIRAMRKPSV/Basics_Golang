package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	Course_Id   string `json:"course_id`
	Course_Name string `json:"course_name`
	Course_cost int    `json:"price"`
	Author      *Author
}

type Author struct {
	Author_Name string `json:"author_name"`
	Experience  int    `json:"experience"`
}

var courses []Course

func (c *Course) isempty() bool {
	return c.Course_Id == "" && c.Course_Name == ""
}

func main() {
	r := mux.NewRouter()
	courses = append(courses, Course{
		Course_Id:   "3",
		Course_Name: "python",
		Course_cost: 2000,
		Author: &Author{
			Author_Name: "sri",
			Experience:  2}})
	r.HandleFunc("/", serverh).Methods("GET")
	r.HandleFunc("/courses", getdata).Methods("GET")
	r.HandleFunc("/course", create).Methods("POST")
	r.HandleFunc("/course/{id}", deleteOne).Methods("DELETE")
	r.HandleFunc("/course/{id}", update).Methods("POST")

	log.Fatal(http.ListenAndServe(":4000", r))

}

func serverh(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to my website</h1>"))
}

func getdata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func create(w http.ResponseWriter, r *http.Request) {
	var course Course
	w.Header().Set("content-type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&course)
	courses = append(courses, course)
	fmt.Println("the data has been entered")
}

func deleteOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.Course_Id == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}

	}

}

func update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for index, course := range courses {
		if course.Course_Id == params["id"] {
			var course Course
			courses = append(courses[index:], courses[:index+1]...)
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.Course_Id = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return

		}
	}

}
