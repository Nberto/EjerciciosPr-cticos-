// ejercicio2
package main

import (
	"fmt"
	"strconv"
	"time"
)

/*Agregar o restar uno o varios días a una fecha.*/
func main() {
	fmt.Println("Calcular dias entre fechas y agregar un dia mas a la fecha final!")
	// FORMATO DE FECHA ---->> DD/MM/YYYY
	// ejecucion de la funcion para obtener los dias entre las fechas establecidas
	f1 := "03-10-2016"      // asigno la var f1 y dentro de ella la cadena q contiene la fecha inicial
	f2 := "06-12-2016"      // asigno la var f2 y dentro de ella la cadena q contiene la fecha final
	diasentrefechas(f1, f2) // llamo a la funcion y le paso los parametros tipo cadena q contienen las variables f1,f2
}

// declaro una funcion con los parametros tipo cadena f1, f2 para recibir los valores a ejecutar en el main y q me devuelva un valor en entero q es el numero de dias

func diasentrefechas(f1, f2 string) int {
	const shortForm = "02-01-2006" // DD/MM/YYYY declaro la constante q me va a permitir modificar el formato de fecha a mi gusto
	// una vez se captura los datos por f1,f2 se realiza el parseo de cadena a time pasando como referencia la constante del formato de mi eleccion
	fechainicio, _ := time.Parse(shortForm, f1) // declaro la var fechainicio,fechafin y dentro de ella invoco la funcion time añadiendo el componente parse
	fechafin, _ := time.Parse(shortForm, f2)    // y pasando la constante de formato y el valor a convertir en time

	// para agregar un dia mas declaro la var adicionar
	adicionar := fechafin.Add(time.Hour * 24)   // para adicionar un dia a un fecha se debe llama la funcion add de time y dentro llamar la func time junto con el componente de horas y lo multiplico por 24 q es igual a 1 dia todo esto dentro de la var adicionar q va a contener la fecha del dia agregado
	duraciontotal := adicionar.Sub(fechainicio) // para calcular los dias tomo las horas como ref esta vez tomando la nueva fecha final y llamado al componente sub de time pasandole el parametro q contiene la fechainicial parseada anteriormente
	resultdias := duraciontotal.Hours() / 24    //para calcular los dias entre estas fechas declaro la varresultdias dentro de ella llamo la var duracion total q contiene las horas y llamo al metodo de time horas y lo divido sobre 24 q es igual a un dia el resultdo de esto sera la cantidad de dias entre las fechas
	//para establecer el formato constante para cuando se imprima debemo de llamar a la componente format de la funcion time q se encargara de hacer la conversion a la establecida en la constante
	formatf1 := fechainicio.Format(shortForm) //declaro la var formatf1 q va a contener el resultado y llamo a la var fechainicio q tiene la fecha parseada y luego al componente format de la funcion time y le paso la constante o la var q tiene el formato predeterminado
	formatf2 := adicionar.Format(shortForm)
	formatf3 := fechafin.Format(shortForm)
	fmt.Println("Fecha Inicial 1 -> ", formatf1) //imprimo la fecha1 con el formato establecido
	fmt.Println("Fecha Final 2 -> ", formatf3)   //imprimo la fecha2 con el formato establecido
	fmt.Println("Fecha Final 2 + dia Adicional  -> ", formatf2)
	fmt.Println("Dias Totales :", resultdias)             //imprimo la cantidad de dias entre la fecha inicial y el dia agregado a la fecha final
	conver := strconv.FormatFloat(resultdias, 'f', 0, 32) // convierto el resultado de los dias de la var result q es tipo float64 a tipo cadena
	conver2, _ := strconv.Atoi(conver)                    // convierto el valor tipo cadena almacenado en la var conver y lo paso a tipo entero
	return conver2                                        // y retorno como resultado de esta funcion un valor tipó entero q esta almacenado en la var conver2
}
