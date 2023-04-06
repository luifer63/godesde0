package arreglos_slices

import "fmt"

var TablaSlice[]int = []int{2,5,7}

var arreglo[10]int = [10]int{6,55,88,99,66,4,563,14}

func MuestroSlices(){

	fmt.Println(TablaSlice)
	porcion := arreglo[3:]   // slice creado desde un vector, desde la posicion 3
	porcion2 := arreglo[:5]  // slice creado desde un vector, desde la posicion 0 hasta la 5
	porcion3 := arreglo[6:7] // slice creado desde un vector, desde la posicion 6 hasta la 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad(){
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i:=0; i<200; i++{
		nums = append(nums, i)
	}

	fmt.Printf("\nLargo %d, capacidad %d", len(nums), cap(nums))
}