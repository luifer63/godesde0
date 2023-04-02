package main

import (
	"fmt"
	"runtime"
	//"github.com/luifer63/godesde0/variables"
	"github.com/luifer63/godesde0/ejercicios"
)

// funcion principal de go
func main(){
	// estado, texto := variables.ConviertoaTexto(1588)
	// fmt.Println(estado)
	// fmt.Println(texto)
	if os := runtime.GOOS; os == "linux" || os == "OX X."{
		fmt.Println("Esto no es windows es ", os)
	} else {
		fmt.Println("Esto es windows")
	}


	switch os:= runtime.GOOS; os{
	case "linux":
		fmt.Println("Esto es linux")
	case "darwin":
		fmt.Println("Esto es darwin")
	default:
		fmt.Printf("%s \n", os)	
	}

	entero, texto := ejercicios.DevolverEntero("101")
	fmt.Println(entero)
	fmt.Println(texto)
}
