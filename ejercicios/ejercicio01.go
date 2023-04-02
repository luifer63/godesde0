package ejercicios

import (
	"strconv"
)

func DevolverEntero(numero string)(int, string){
	entero, error := strconv.Atoi(numero)
	if error != nil{
		return entero, "Hubo un error " + error.Error()
	}
	if  entero >= 100{
		return entero, "Es mayor a 100"
	}else{
		return entero, "Es menor a 100"
	}
}