package ejercicios

import (
	//"fmt"
	"strconv"
)

func DevolverEntero(numero string)(int, string){

	if entero, error := strconv.Atoi(numero); entero >= 100{
		return entero, "Es mayor a 100"
	}else if entero < 100{
		return entero, "Es menor a 100"
	}else{
		return entero, error.Error()
	} 
}