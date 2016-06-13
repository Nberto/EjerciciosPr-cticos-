// ejercicio5_test
package main

import (
	"testing"
)

func TestEjercicio5(t *testing.T) {
	eva := "hola mayus now!"       // var q contiene datos tipo cadena a ejecutar en la funcion
	exp := "HOLA MAYUS NOW!"       // var q contiene el resultado esperando a la ejecucion de la funcion
	resultado := obtenermayus(eva) // var resultado contiene la funcion a evaluar y la palabra q ejecutamos en el proceso
	if resultado != exp {          // declaramos una condicion para determinar si el resultado de la funcion es diferente al esperado
		t.Fatalf("error de ejecucion %V valor esperado -> %v ", resultado, exp) // si es diferente se imprime un mensaje visualizando el resultado de la ejecucion y el valor esperado
	}
}
