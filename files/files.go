package files

import (
	
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/luifer63/godesde0/ejercicios"
)

var filename string = "./files/txt/tablasDeMultiplicar.txt"

func GrabaTabla(){
	var texto string = ejercicios.TablaDelNumero()
	archivo, err := os.Create(filename)
	if err != nil{
		fmt.Println("Error al crear el archivo "+ err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla(){
	var texto string = ejercicios.TablaDelNumero()
	if !Append(filename, texto){
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filen string, texto string) bool{
	arch, err := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND, 0644)
	if err != nil{
		fmt.Println("Error durante el append " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)

	if err != nil{
		fmt.Println("Error durante el WriteString " + err.Error())
		return false
	}

	arch.Close()
	return true
}


func LeoArchivo(){
	archivo, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println("Error al leer archivo " + err.Error())
		return
	}

	fmt.Println(string(archivo))
}

func LeoArchivo2(){
	archivo, err := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND, 0644)
	if err != nil{
		fmt.Println("Error al leer archivo " + err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}

	archivo.Close()
}