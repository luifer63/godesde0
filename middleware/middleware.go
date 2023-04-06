package middleware

import "fmt"

func sumar(a, b int) int{
	return a + b
}

func restar(a, b int) int{
	return a - b
}


func multiplicar(a, b int) int{
	return a * b
}

func MiMiddleWare(){
	fmt.Println("Inicio")

	result := operacioensMidd(sumar)(2,3)
	fmt.Println(result)
	result = operacioensMidd(restar)(2,3)
	fmt.Println(result)
	result = operacioensMidd(multiplicar)(2,3)
	fmt.Println(result)
}

func operacioensMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int{
		fmt.Println("Inicio de operación")
		return f(a, b)
	}
}