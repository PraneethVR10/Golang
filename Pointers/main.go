package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func modifyValue(x *int) {
	*x = 20
}

func main()  {
	var a int 
	
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	numModify, err  := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		a =int(numModify)

		fmt.Println("Before modification:" ,a)

		modifyValue(&a)

		fmt.Println("After modification:", a)
	}


	

	
	
}
