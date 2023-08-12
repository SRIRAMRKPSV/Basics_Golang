package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// TO CREATE A DB STRUCTURE
type Course struct {
	Course_name string
	Course_id   string
	Rate        int `json:"price"`
	Author      *Author
}

type Author struct {
	Name string `json:"author_name"`
	Site string `json:"website"`
}

var courses []Course

func (c *Course) IsEmpty() bool {

	//return c.Course_name == "" && c.Course_id == ""
	return c.Course_name == ""
}

func main() {
	fmt.Println("welcome to the Program")

	r := mux.NewRouter()

	courses = append(courses, Course{Course_name: "python", Course_id: "2", Rate: 499, Author: &Author{Name: "sriram", Site: "sriram.ai"}})
	courses = append(courses, Course{Course_name: "full_stack", Course_id: "4", Rate: 1499, Author: &Author{Name: "sri", Site: "myweb.ai"}})
	courses = append(courses, Course{Course_name: "Golang", Course_id: "5", Rate: 299, Author: &Author{Name: "ram", Site: "lco.ai"}})

	// routing
	r.HandleFunc("/", responses).Methods("GET")
	r.HandleFunc("/courses", final).Methods("GET")
	r.HandleFunc("/course/{id}", add).Methods("GET")
	r.HandleFunc("/course", createcourse).Methods("POST")
	r.HandleFunc("/course/{id}", updatecourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deletecourse).Methods("DELETE")

	// to connect in local host

	log.Fatal(http.ListenAndServe(":3000", r))

}

// TO DISPALY A HTML

func responses(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1><marquee>Hello Every One<marquee></h1>"))

}

// TO SHOW THE COURSE IN DATABASE

func final(w http.ResponseWriter, r *http.Request) {
	fmt.Println("welcome to the first api")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

// TO ADD A COURSE

func add(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get the request")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.Course_id == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	json.NewEncoder(w).Encode("NO course is found")
	return

}

//  TO CREATE A COURSE

func createcourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get the request")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("there is no course")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("add the new course, no data inside json")
		return
	}
	rand.Seed(time.Now().UnixNano())
	course.Course_id = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

// to update the data
func updatecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update the values")
	w.Header().Set("Conent-Type", "application/json")

	params := mux.Vars(r)

	// loop,id , remove,add

	for index, course := range courses {
		if course.Course_id == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.Course_id = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}

	}
}

// to delete the course
func deletecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one of the values")
	w.Header().Set("Conent-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.Course_id == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

}
