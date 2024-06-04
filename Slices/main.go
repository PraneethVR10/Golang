package main

import (
	"fmt"
)

func main() {
	 var numbers = []int{1,2,3,4,5}
    fmt.Println(numbers)
var index int = 1
	numbers = append(numbers[:index],numbers[index+1:]...)
	fmt.Println(numbers)

  languages := make(map[int]string)

  languages[1] = "python"
  languages[2] = "JS"
  languages[3] = "Ruby"  
  fmt.Println(languages)
  fmt.Println(languages[2])
  delete(languages, 3)

  for err, value := range languages {
	fmt.Println(err, value)
  }

}

