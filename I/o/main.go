package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the pizza rating")

	input, _:= reader.ReadString('\n')
	fmt.Println("The rating you have given is :", input)
	
	
}