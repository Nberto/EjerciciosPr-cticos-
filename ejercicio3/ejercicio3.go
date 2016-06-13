// ejercicio3
package main

import (
	"fmt"
	//"strconv"
	"strings"
)

/* Crear una función que rellene por la izquierda una cadena de caracteres
hasta una longitud dada, con un carácter específico, por ejemplo: la
cadena(“GOLANG”), longitud dada (10), caracter para rellenar (“A”).
Resultado final “AAAAGOLANG”.*/

func main() {
	//addpalabra() // ejecucion de la funcion addpalabra
	var str = "GOLANG"
	fmt.Println(rellenopalabra(str, "a", 10))
	fmt.Println(rellenoContador(str, "A", 10))
}

// creo la funcion rellenopalabra y dentro de ella invoco los parametros del metodo Repeat de la func strings
func rellenopalabra(s string, padStr string, plen int) string { // asgino variables s,padStr tipo cadena q albergara la palabra y el valor a agregar a la palabra ademas de crear la var plen q va a contener el numero de veces q se repetira el valor a agregar a la palabra
	return strings.Repeat(padStr, plen) + s // retorno el valor de la misma en tipo cadena y asgino los parametros de salida de la misma para q se ejecuten de la siguiente manera
}

// creo la funcion rellenocontador la cual va realizar la operacion de calcular la longitud de la palabra y agregar el valor cadena a la palabra orginal sin que se pase de la longitud establecida
func rellenoContador(s string, padStr string, overrall int) string { // dentro de la func declaro las variable s,padStr tipo cadena q albergara la palabra y el valor a agregar a la palabra ademas de crear la var plen q va a contener el numero de veces q se repetira el valor a agregar a la palabra
	var padcountInt int                                        // inicializo la var padcountInt en tipo entero q me permitira realizar la cuenta de el rango establecido por las palabras y q no se pase de la longitud establecida
	padcountInt = 1 + ((overrall - len(padStr)) / len(padStr)) // dentro de var PadcountInt declaramos valor de 1 me sume y dentro de los ()  declaramos q longitud sea menor al rango de la palabra y dividimos sobre el mismo rango con el fin de obtener el resultado q nos permita determinar el valor cadena a agregar a la palabra
	var retStr = strings.Repeat(padStr, padcountInt) + s       // declaramos a var reStrs q realizara la operacion de agregar o repetir el valor cadena en la palabra hasta la longitud establecida y determinada por la var padcountInt
	return retStr[(len(retStr) - overrall):]                   // retornamos la palabra y calculamos si el rango es el estableido para q imprima solo la longitud establecida
}

/*
// creo la funcion addplabra y dentro de ella
func addpalabra() {
	n := 1                  // inicializo la var n = 1
	conv := strconv.Itoa(n) //luego declaro la var conv dentro de ella invoco al componente strconv de la funcion strconv y paso entero a cadena
	add := conv + "GOLANG"  // declaro la var add y dentro de ella se agrega la var conv q tiene el valor caden a agregar mas la palabra q va a tener el nuevo valor
	fmt.Println(add)        // imprimop la var add q contiene el resultado de agregar la cadena 1 a la cadena golang
}
*/
