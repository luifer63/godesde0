package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func TablaDelNumero() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Ingrese número: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("El dato ingresado es incorrecto " + err.Error())
			} else {
				for i := 1; i <= 10; i++ {
					fmt.Printf("%d x %d = %d \n", numero, i, numero*i)
				}
				break
			}
		}
	}
}
