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

// Model  for couse - file
type Course struct {
	CourseId   string  `json:"courseid"`
	CourseName string  `json:"coursename"`
	Price      int     `json:"price"`
	Author     *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake database
var courses []Course

// middlewarte, helper - file

func IsEmpty(c *Course) bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("Buliding an Api")
	r := mux.NewRouter()
	courses = append(courses, Course{CourseId: "2",
		CourseName: "Golang",
		Price:      0,
		Author:     &Author{FullName: "Hitesh", Website: "chaiaurcode"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getallCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOnecourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

//controller - file

//serverHome controller

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("servingHome")
	w.Write([]byte("<h1>The response by seveHome controller</h1>"))
}

func getallCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting all the courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOnecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request

	params := mux.Vars(r)

	// loop through courses, find matching course

	for _, val := range courses {
		if val.CourseId == params["id"] {
			json.NewEncoder(w).Encode(val)
			return
		}
	}
	json.NewEncoder(w).Encode("No such courses found")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating the new course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please fill the required field")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if IsEmpty(&course) {
		json.NewEncoder(w).Encode("Please fill the required field")
		return
	}

	for _, val := range courses {
		if course.CourseName == val.CourseName {
			json.NewEncoder(w).Encode("Sorry, course is already present")
			return
		}
	}
	// generate uniqueid
	// append the course in the courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updating the course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	if params["id"] == "" {
		json.NewEncoder(w).Encode("please provide the id")
		return
	}

	for index, val := range courses {
		if params["id"] == val.CourseId {
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses[index] = course
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("sorry, no such course id is found")
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleting the course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	if params["id"] == "" {
		json.NewEncoder(w).Encode("please provide the id")
		return
	}

	for index, val := range courses {
		if params["id"] == val.CourseId {
			//remove the value
			courses = append(courses[:index], courses[(index+1):]...)
			json.NewEncoder(w).Encode("Course succesfully deleted")
			return
		}
	}

	json.NewEncoder(w).Encode("sorry, no such course id is found")

}
