package main

import "fmt"

/*

func helloworld() string {

	return "Hola Mundo"
}
*/

func estudiante() (string, string) {
	return "Andres Fernandez", "Ingenieria de Software"
}

func agregarDatos(nombre string, apellido string, edad int) map[string]interface{} {
	var datos = map[string]interface{}{
		"nombre":   nombre,
		"apellido": apellido,
		"edad":     edad,
	}

	datos["Carrera"] = "Ingenieria de software"

	return datos
	//return datos["nombre"].(string), datos["apellido"].(string), datos["edad"].(int)
}

func main() {

	//holamundo := helloworld()
	//var nombre string = " Andres Fernandez"

	var datos string = "Estos son los datos del estudiante, \n Nombre: %s \n Carrera: %s"
	nombre, carrera := estudiante()

	var slice = []int{1, 2, 3, 4}
	fmt.Printf(datos, nombre, carrera)
	fmt.Printf("Slice %v", slice)

	//var nombreData, apellidoData, edadData = agregarDatos("Andres", "Fernandez", 22)

	//fmt.Printf("\n Nombre: %s, Apellido: %s, Edad: %d", nombreData, apellidoData, edadData)

	var datosMap map[string]interface{} = agregarDatos("Andres", "Fernandez", 22)

	fmt.Print("\nListo para recorrer MAP \n\n")

	var nuevaLista []interface{}

	for k, v := range datosMap {

		fmt.Printf("Clave: %s Valor: %v \n", k, v)
		nuevaLista = append(nuevaLista, v)
	}

	fmt.Print("\n Lista: \n")
	fmt.Print(nuevaLista)

}
