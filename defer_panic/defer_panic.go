package defer_panic

import (
	//"os"
	"log"
	"fmt"

)


func VemosDefer(){
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es un ultimo mensaje")
	fmt.Println("Este es un segundo mensaje")
}

func EjemploPanic(){
	defer func(){
		reco := recover()

		if reco != nil {
			log.Fatalf("Ocurrio un errror que generó PANIC \n %v", reco)
		}
	}()
	a:=1
	if a == 1 {
		panic("Se encontró el valor 1")
	}
}