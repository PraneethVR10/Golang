package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)



type Course struct {

	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author     *Author `json:"author"` 
}

type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

var courses []Course

func(c *Course) IsEmpty() bool {
return c.CourseId == ""  && c.CourseName == "" 
}
func main() {
	
}

func serveHome (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Praneeth</h1>"))

}

func getAllCoureses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")

w.Header().Set("Content-Type","Application/json")
json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type","application/json")

      params := mux.Vars(r)
	  fmt.Printf("The type of Params is :%T/n", params)

	   for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course) 
			return 
		}
	   }
	   json.NewEncoder(w).Encode("No course found with given id")
	   return 

}