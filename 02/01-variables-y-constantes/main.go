package main

import "fmt"

func main() {
	nombre, apellido := "Tu nombre", "Tu apellido"
	nombre = "Carlos"           // nueva asignación de valor
	nombre, edad := "Pedro", 30 // nueva asignación de nombre con := , generando una nueva
	fmt.Println(nombre, apellido, edad)
}
