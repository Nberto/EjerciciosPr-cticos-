// ejercicio4_test
package main

import (
	"testing"
)

func TestEjercicio4(t *testing.T) {
	exp := "GOO"                              // declaro la expectativa al resultado de la funcion a evaluar
	str := "GOOGLE"                           // declaro la frase a evaluar para posteriormente disminuir su tamaño de acuerdo a la cantidad de letras a restar
	resultado := extraerpalabras(str, "a", 3) // declaro la var resultado y dentro de ella paso la funcion a evaluar, dentro del parentesis paso la palabra a calcular , la referencia a reducir es decir la letra en su caso cualquiera y por ultimo la cantidad a reducir
	if resultado != exp {                     // determno la condicion q me permita evaluar si el resultado es igual o diferente al esperado
		t.Fatalf("error al ejecutar %v valor esperado -> %v", resultado, exp) // si es diferenteimprime un mensaje con el valor obtenido por la funcion y el valor esperado como expectativa correcta
	}
}
