package ejer_interfaces

import (
	"fmt"

	"github.com/luifer63/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	hu.EstaVivo()
	fmt.Printf("Soy un/a %s y estoy respirando \n y vivo %t\n", hu.Sexo(), hu.EstaVivo())
}