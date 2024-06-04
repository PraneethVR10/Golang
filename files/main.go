package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Create("./sonu.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString("Hello, This is Sonu")

	if err != nil {
		panic(err)
	}

	content , err := ioutil.ReadFile("./sonu.txt")

	if err != nil {
		panic(err)
	}else {
		fmt.Println(string(content))
	}



  filepath  := "./sonu.txt"
  if _, err := os.Stat(filepath); os.IsNotExist(err) {
	fmt.Println("The file does not exist")
} else {
	fmt.Println("The file exists")
}
  err = os.Remove(filepath)
  if err != nil {
	panic(err)
   } else {
	fmt.Println("This file has been deleted")
   }
}




