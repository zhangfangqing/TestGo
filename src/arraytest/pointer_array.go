package main

import "fmt"


func f(arr [5]int){
	fmt.Println(arr)
}

func fp(arr *[5]int){
	fmt.Println(arr)
}

func main(){
	var arr [5]int
	f(arr)
	fp(&arr)
}