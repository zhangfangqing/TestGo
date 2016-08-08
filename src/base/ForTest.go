package main

import (
	"fmt"
)

func main() {

	var str string = "aabbd"

	for i := 0 ; i < len(str) ; i++  {
		fmt.Printf("character on position %d is : %c \n" , i , str[i])
	}

	fmt.Println("==========================================")

	for idex , val := range str{
		fmt.Printf("character on position %d is : %c \n" , idex , val)
	}
}
