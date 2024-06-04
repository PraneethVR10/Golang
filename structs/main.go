package main

import "fmt"

func main() {
 
	Praneeth := User{"Praneeth","praneethvr03@gmail.com",true,23}
	fmt.Println(Praneeth)
	fmt.Printf("%+v\n",Praneeth)
}
	type User struct {
		Name string	
		Email string
		Status  bool
		Age  int
	}
