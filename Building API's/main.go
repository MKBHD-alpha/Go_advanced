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

// Model for Courses - file

type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course-name"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB

var courses []Course

// middleware, helpers -file

// (c *Course) -> because were passing a part of the struct
func (c *Course) IsEmpty() bool {
	//	return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == "" //manually providing the CourseID
}

func main() {
	fmt.Println("API - LCO")
	r := mux.NewRouter() //creating  a new router

	//seeding (using fake data)
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})

	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listening to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers -file (handling situations)

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to API Building </h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Showing All Course")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) //Encode(courses)  sending API Json Response

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)
	fmt.Println(params)

	//loop through courses, find matching id and return response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found with  given ID")
	return

}

// Adding Data
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("One Course Added")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some Data")
	}
	// body -> {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No Data Inside Json")
		return
	}
	//DONE:checking duplicate
	for _, name := range courses {
		if name.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course name Already Exists")
			return

		}
	}
	//generating unique id and converting to string
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

// updating course
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("One Course Updated")
	w.Header().Set("Content-Type", "application/json")
	//first - grab id from request body
	params := mux.Vars(r)

	//loop, id, remove, add with  id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return

			//DONE: Send a response when id is not found
		}

	}
	json.NewEncoder(w).Encode("Unable to find Course with given ID")
	return

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("One Course Deleted")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course Deleted Successfully")
			break

		}
		
	}

}
