package main

import "fmt"

func main() {
	var praneeth_age int = 17
	var result string = "he is eligible to drink"

	if praneeth_age>18 {
		fmt.Println(result)
	} else if praneeth_age ==18 {
		fmt.Println(result)
		
	} else {
		fmt.Println("He is not eligible to drink")
	}

	if new_age := 10; new_age>5 {

		fmt.Println("you are awesome")
	} else {
		fmt.Println("you are not awesome")
	}
}