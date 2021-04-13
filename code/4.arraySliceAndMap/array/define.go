package main
import "fmt"

func main(){
	var array1 [5]int
	array2 := [5]int{10,20,30,40,50}
	array3 := [...]int{10,20,30,40,50}
	array4 := [...]int{0:0,1:10,2:20,3:0,4:0}

	for i, array := range [4][5]int{array1, array2, array3, array4}{
		fmt.Printf("%d definition in array %d:\n", i, &array)
		for _, item := range array{
			fmt.Print(item)
			fmt.Println()
		}

	}
}
