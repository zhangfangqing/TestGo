package main

import "fmt"

func main(){

	var myarr [5]int

	for i := 0 ; i < len(myarr) ; i++  {
		myarr[i] = i*2
	}

	for i := 0 ; i < len(myarr) ; i++  {
		fmt.Println(myarr[i])
	}


}