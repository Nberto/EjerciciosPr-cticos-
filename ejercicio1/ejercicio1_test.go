// ejercicio1_test
package main

import (
	"testing"
)

func TestEjercicio1(t *testing.T) {
	f1 := "03-10-2016"                   // asigno la var f1 y dentro de ella la cadena q contiene la fecha inicial
	f2 := "03-11-2016"                   // asigno la var f2 y dentro de ella la cadena q contiene la fecha final
	expDias := 31                        // declaro el valor a esperar como resultado de la ejecucion de la funcion
	resultado := diasentrefechas(f1, f2) // declaro una var resultado q va a contener la funcion a ejecutar junto con las var q tiene los datos a calcular
	if resultado != expDias {            // declaro la condicion pasando las variables y determinando q si el resultado de la funcion es diferente al valor esperado q declare un error
		t.Fatalf("error en numero de dias calculados %v con relación a los días esperados %v", resultado, expDias) // dicho error se imprime llamando a la var q contiene la respuesta de la funcion testion y el metodo faltaf
	}
}
