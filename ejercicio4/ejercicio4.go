// ejercicio4
package main

import (
	"fmt"
	"strings"
)

func main() {
	//declramos la var str q contiene la palabra a reducir
	var str = "GOOGLE"
	fmt.Println(extraerpalabras(str, "A", 3)) // imprimimos el valor a ejecutar en la cadena reduciendola de acuerdo a la longitud establecida
}

/* Crear una función Substr, que extraiga de una cadena de caracteres
dada, una parte de esta cadena con una longitud dada. Por ejemplo:
Cadena dada (“GOOGLE”), longitud 3. Resultado “GOO”. */

// creo la funcion extraer palabras y declaro los parametros  con los cuales se va a evaluar la funcion o ejecutar devolviendo una cadena al final
func extraerpalabras(s string, padStr string, overall int) string { // declaramos las var s,padStr tipo cadena q van a contener el valor tipo cadena y el otro la palabra  ademas de overral q es tipo entero con el fin de determinar la longitud de la misma oracion a construir
	padcountInt := 1 + ((overall - len(padStr)) / len(padStr)) // dentro de var PadcountInt declaramos valor de 1 me sume y dentro de los ()  declaramos q longitud sea menor al rango de la palabra y dividimos sobre el mismo rango con el fin de obtener el resultado q nos permita determinar el valor cadena a restar a la palabra
	retStr := s + strings.Repeat(padStr, padcountInt)          //declaramos a var reStrs q realizara la operacion de agregar o repetir el valor cadena en la palabra hasta la longitud establecida y determinada por la var padcountInt
	return retStr[:overall]                                    //retornamos la palabra y calculamos si el rango es el estableido para q imprima solo la longitud establecida

}
