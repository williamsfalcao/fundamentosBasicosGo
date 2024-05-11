package main

import "fmt"

func main() {

	fmt.Println("Array e Slice")

	var array0 [5]string
	array0[0] = "Posicao 1"
	fmt.Println(array0)

	array1 := [5]string{"Posicao 1", "Posicao 2", "Posicao 3", "Posicao 4", "Posicao 5"}
	fmt.Println(array1)

	slice := []int{10, 11, 12, 13, 14, 15, 16}
	fmt.Println(slice)

	slice = append(slice, 19)
	fmt.Println(slice)

	//Array internos
	slice1 := make([]float32, 5, 6)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1 = append(slice1, 43)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1 = append(slice1, 87)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

}
