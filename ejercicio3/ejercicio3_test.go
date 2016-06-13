// ejercicio3_test
package main

import (
	"testing"
)

func TestEjercicio3(t *testing.T) {
	exp1 := "aaaaaaaaaaGOLANG" // declaro la expectativa al resultado de la funcion la cual sera la respuesta correcta
	exp2 := "AAAAGOLANG"
	str := "GOLANG"                            // declaro la palabra a rellenar con mas datos tipo cadena
	resultado1 := rellenopalabra(str, "a", 10) // declaro la var resultado1 y dentro llamo al metodo a evaluar pasando los parametros q recibe para evaluar q son la var str q contiene la palabra , la letra a agregar , la cantidad a agregar
	resultado2 := rellenoContador(str, "A", 10)
	// declaro la condicion para determinar si el resultado de la funcion es el esperado a la var q contiene la respuesta correcta
	if resultado1 != exp1 {
		t.Fatalf("error de ejecucion %v , valor esperado ->", resultado1, exp1) // en caso de no ser correcto imprime un mensaje visualizando el resultado de la funcion y la expectativa a ese resultado
	}
	if resultado2 != exp2 {
		t.Fatalf("error de ejecucion %v , valor esperado ->", resultado2, exp2)
	}
}
