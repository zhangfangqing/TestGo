package main

import "fmt"

func add(a,b,c int) int {
	return a + b + c
}

func main() {
	resp :=  add(1,2,5)
	fmt.Println(resp)

}