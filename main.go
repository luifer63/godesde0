package main

import (
	"fmt"

	"github.com/luifer63/godesde0/variables"
)

// funcion principal de go
func main(){
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

}
