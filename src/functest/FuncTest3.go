package main

import (
	"fmt"
)

func min(num ...int) (resp int) {
	if len(num) == 0 {
		return 0
	}
	min := num[0]
	for _ , val := range num{
		if min > val {
			min = val
		}
	}
	return  min
}

func main() {
	aa := min(11,2,5,3)
	fmt.Println(aa)

	bb :=  min(3,4,72,1,7,8)
	fmt.Println(bb)

}
