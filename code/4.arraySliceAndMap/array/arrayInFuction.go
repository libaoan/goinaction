package main

import "fmt"

func main(){
	var array1 [1e6]int
	foo(array1)
}

func foo(array [1e6]int){
	fmt.Println(&array, len(array))
}

