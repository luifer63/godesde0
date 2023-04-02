package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

func TablaDelNumero() string {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Ingrese número: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("El dato ingresado es incorrecto " + err.Error())
			} else {
				for i := 1; i <= 10; i++ {
					texto += fmt.Sprintf("%d x %d = %d \n", numero, i, numero*i)
				}
				break
			}
		}
	}
	return texto
}
