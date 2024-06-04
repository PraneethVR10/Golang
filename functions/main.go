package main

import "fmt"

func main() {
	fmt.Println("Hi, Below is the example of a function")
	praneeth()
	result,message  := sonu(5, 6)
	fmt.Println("Result:", result,message)
	defer fmt.Println("Hey, How are you doing")
}

func praneeth() {
	fmt.Println("This is the actual function")
}

func sonu(val_1 int, val_2 int) (int,string) {
	return val_1 + val_2, "Hi is the sum"
}
