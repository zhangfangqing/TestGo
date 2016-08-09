package main

import (
	"fmt"
)

func Add(a int , b int) {
	fmt.Println(a + b)
}

func Mutiful(a int , b int) {
	fmt.Println(a * b)
}

func Dev(a int , b int) {
	fmt.Println(a / b)
}

func CallBack(a int , f func(int ,int)) 	{
	f(a , 2)
}

func main()  {
	CallBack(1 , Add)
	CallBack(3 , Mutiful)
	CallBack(10 , Dev)
}

