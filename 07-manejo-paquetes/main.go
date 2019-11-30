package main

import (
	"fmt"

	"github.com/LeoPindjs/primeros-pasos-en-golang/07-manejo-paquetes/saludar"
	"github.com/LeoPindjs/primeros-pasos-en-golang/07-manejo-paquetes/sumar"
)

func main() {
	s := saludar.Saludar("Leo")
	fmt.Println(s)
	t := sumar.Sumar(14, 30.5, 20, 35.5, 5.8)
	fmt.Printf("La suma de los dato es %f\n", t)

}
