// Ejercicio1
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	//const shortForm = "01-02-2006" // DD/MM/YYYY
	/*t, _ := time.Parse(shortForm, "10-03-2016")
	t2, _ := time.Parse(shortForm, "11-03-2016")
	fmt.Println("fecha inicial ->", t)
	fmt.Println("fecha final ->", t2)
	duracion := t2.Sub(t)
	fmt.Println("numero de horas :", duracion)
	result := duracion.Hours() / 24
	fmt.Println("numero de dias :", result)
	adicionar := t2.Add(time.Hour * 24)
	result2 := t2.Format(shortForm)
	fmt.Println("dias adicionales :", adicionar)
	fmt.Println(result2)*/
	// FORMATO DE FECHA ---->> DD/MM/YYYY
	// ejecucion de la funcion para obtener los dias entre las fechas establecidas
	f1 := "03-10-2016"      // asigno la var f1 y dentro de ella la cadena q contiene la fecha inicial
	f2 := "03-11-2016"      // asigno la var f2 y dentro de ella la cadena q contiene la fecha final
	diasentrefechas(f1, f2) // llamo a la funcion y le paso los parametros tipo cadena q contienen las variables f1,f2
}

//1-02-2016
/*Calcular la cantidad de días entre dos fechas. Tener en cuenta que las
fechas deben estar sin tiempo, es decir, deben tener el formato
(DD/MM/YYYY).*/

// declaro una funcion con los parametros tipo cadena f1, f2 para recibir los valores a ejecutar en el main y q me devuelva un valor en entero q es el numero de dias
func diasentrefechas(f1, f2 string) int {
	const shortForm = "02-01-2006" // DD/MM/YYYY declaro la constante q me va a permitir modificar el formato de fecha a mi gusto
	// una vez se captura los datos por f1,f2 se realiza el parseo de cadena a time pasando como referencia la constante del formato de mi eleccion
	fechainicio, _ := time.Parse(shortForm, f1) // declaro la var fechainicio,fechafin y dentro de ella invoco la funcion time añadiendo el componente parse
	fechafin, _ := time.Parse(shortForm, f2)    // y pasando la constante de formato y el valor a convertir en time
	//para determinar la cantidad de horas entre las fechas invoco el metodo de time llamando a la var fechafin y añadiendo el componente sub q me permitira
	duracion := fechafin.Sub(fechainicio)     // saber el numero de horas entre las fechas
	result := duracion.Hours() / 24           // para determinar la cantidad de dias creamos la var result q contiene la var duracion la cual alberga el tiempo en horas luego se llama al componente horas y los divide sobre 24 q es igual a 1 dia
	formatf1 := fechainicio.Format(shortForm) //declaro la var formatf1 q va a contener el resultado y llamo a la var fechainicio q tiene la fecha parseada y luego al componente format de la funcion time y le paso la constante o la var q tiene el formato predeterminado
	formatf2 := fechafin.Format(shortForm)
	fmt.Println("Fecha Inicial 1 ->", formatf1) // imprimo la fecha q hizo el parseo de string a time
	fmt.Println("Fecha Final 2 ->", formatf2)   // imprimo la fecha q hizo el parseo de string a time
	//fmt.Println(" Numero de Horas :", duracion) // imprimo la duracion del tiempo entre las fechas
	fmt.Println("Dias entre Fechas iniciales :", result) // imprimo la cantidad de dias entre las primeras fechas establecidas
	conver := strconv.FormatFloat(result, 'f', 0, 32)    // convierto el resultado de los dias de la var result q es tipo float64 a tipo cadena
	conver2, _ := strconv.Atoi(conver)                   // convierto el valor tipo cadena almacenado en la var conver y lo paso a tipo entero
	return conver2                                       // y retorno como resultado de esta funcion un valor tipó entero q esta almacenado en la var conver2
}
