package main

import (
	"encoding/json"
	"fmt"
)


type course struct {
	Name string `json:"coursename"`
	Price int
	Platform string `json:"website"`
	Password string  `json:"-"`
	Tags []string  `json:"tags"`
} 

func main() {
	EncodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJs Bootcamp",299,"Learncodeonline.in","abc123", []string{"webdev", "js"}}, 
	}
	finalJson, err  := json.MarshalIndent(lcoCourses,"","\t") 

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n",finalJson)
}