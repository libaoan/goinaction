package main

import "fmt"

func main(){
	var array1 [4][2]int
	array2 := [4][2]int{{10,11}, {20,21},{30,31},{40,41}}
	array3 := [4][2]int{1:{20,21}, 3:{40,41}}
	array4 := [4][2]int{1:{0:20}, 3:{1:41}}
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array4)


	var array5 [2]int = array4[1]
	fmt.Println(array5)
}
