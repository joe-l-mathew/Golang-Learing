package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice string  `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

//middleware for checking empty course id or empty name

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

//temperary DB

var courses []Course = []Course{}

func main() {
	SeedDummyData()
	g := mux.NewRouter()

	g.HandleFunc("/", ServeHome).Methods("GET")
	g.HandleFunc("/course", GetAllCourse).Methods("GET")
	g.HandleFunc("/course/{id}", GetOneCourse).Methods("GET")
	g.HandleFunc("/course", CreateOneCourse).Methods("POST")
	g.HandleFunc("/course/{id}", UpdateCourse).Methods("PUT")
	g.HandleFunc("/course/{id}", UpdateCourse).Methods("DELETE")
	http.ListenAndServe(":4000", g)

}

func SeedDummyData() {
	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "JavaScript",
		CoursePrice: "Free",
		Author: &Author{
			FullName: "Author Full Name",
			Website:  "www.google.com",
		},
	}, Course{
		CourseId:    "2",
		CourseName:  "Flutter BLOC",
		CoursePrice: "100",
		Author: &Author{
			FullName: "Felix Angalov",
			Website:  "www.blocflutter.com",
		},
	})

}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is your home for courses</h1>"))
}

func GetAllCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(courses)
}

func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "Application/json")
	params := mux.Vars(r)
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with id")
}

func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "Application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("No body found")
		return
	}
	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Invalid body found")
		return
	}

	rand.Seed(int64(time.Now().Nanosecond()))
	randNumber := rand.Intn(10000)
	course.CourseId = strconv.Itoa(randNumber)
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "Application/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			fmt.Println(params["id"])
			newCourses := append(courses[:index], courses[index+1:]...)
			var updatedCourse Course
			json.NewDecoder(r.Body).Decode(&updatedCourse)
			updatedCourse.CourseId = params["id"]
			courses = append(newCourses, updatedCourse)
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	json.NewEncoder(w).Encode("No course with id found")

}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "Application/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Deleted Succesfully")
			return
		}
	}
	json.NewEncoder(w).Encode("No course with id found")

}
