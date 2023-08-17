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

// model for course -file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fulname string `json:"fullname"`
	Website string `json:"website"`
}

// db
var courses []Course

// middleware,helper-file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("APi Building")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "AndroidDevlopment", CoursePrice: 399, Author: &Author{Fulname: "Sachin", Website: "comingsoon"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Firebase", CoursePrice: 199, Author: &Author{Fulname: "Sachin", Website: "comingsoon"}})

	//routing

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	http.Handle("/", r)
	//listen to port

	log.Fatal(http.ListenAndServe(":4000", r))
}

//controlers -file

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> welcome to API Building part of golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Course")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request

	//loop through courses , find matching id and return response
	vars := mux.Vars(r)
	for _, course := range courses {
		if course.CourseId == vars["id"] {
			json.NewEncoder(w).Encode(course.CourseName)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
	return

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	//what if :body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")
	}

	//what if data -{}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//generate unique id ,string
	//apeend course nto courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	//first grab id from req
	vars := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == vars["id"] {
			courses = append(courses[:index], courses[index+1:]...)
		}
	}
}
