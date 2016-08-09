package main

import "fmt"

func add(a int , b int , c int) (resp1 int , resp2 bool) {
	return a + b + c , true
}

func main() {
	resp , isok :=  add(1,2,3)
	fmt.Println(resp)
	fmt.Println(isok)
	fmt.Println(isok)
}