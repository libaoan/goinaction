package main

import "fmt"

func main(){
	array := [5]int{10,20,30,40,50}
	fmt.Println("array[2] in array", array, array[2])

	array2 := [5]*int{0: new(int), 1: new(int)}
	*array2[0] = 10
	*array2[1] = 20
	for i, j := range array2{
		fmt.Println(i, j)
	}


	var array3 [5]string
	array4 := [5]string{"Red", "Yellow", "Bule", "Black", "Green"}
	array3 = array4
	fmt.Println("after array3 = array4:", &array3, &array4)
	fmt.Printf("%p array3 %d\n", &array3, array3)
	fmt.Printf("%p array4 %d\n", &array4, array4)


	var array5 [3]*string
	array6 := [3]*string{new(string), new(string), new(string)}

	*array6[0] = "Red"
	*array6[1] = "Blue"
	*array6[0] = "Green"

	fmt.Println("before assignment, len of array5", len(array5), array5)
	array5 = array6
	fmt.Println("after assignment, len of array5", len(array5), array5)
}
