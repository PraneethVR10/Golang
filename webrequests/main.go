package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	const url  = "https://www.facebook.com/"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	 fmt.Printf("Response is of Type : %T\n", response)

	  defer response.Body.Close()

	  data, err := ioutil.ReadAll(response.Body)
	   
	  if err != nil {
		panic(err)
	  }
	  content := string(data)

	  fmt.Println(content)

	   
	
}