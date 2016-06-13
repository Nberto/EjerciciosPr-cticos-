// ejercicio5
package main

import (
	"fmt"
	"strings" // funcion q permite manejar valores cadena
)

func main() {
	// llamamos a la func strings y luego al componenteToUpper q se encargara de generar el texto en mayuscula
	//fmt.Println(strings.ToUpper("hola mayus now!")) //impresion del texto en mayus llamando a la funcion
	x := "hola mayus now!"
	fmt.Println(obtenermayus(x))
}

/*Convertir todos los caracteres de una cadena a mayúsculas.*/
// declaro la funcion obtenermayus q me va a capturar un dato tipo string en la var x y me va a devolver un string
func obtenermayus(x string) string {
	return strings.ToUpper(x)
}
